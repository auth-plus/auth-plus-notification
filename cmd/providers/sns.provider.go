package providers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNS struct {
	url   string
	token string
}

func NewSNS() *SNS {
	instance := new(SNS)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *SNS) SendSms(phone string, content string) (bool, error) {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)
	params := &sns.PublishInput{
		Message:     aws.String(content),
		PhoneNumber: aws.String(phone),
	}
	resp, err := svc.Publish(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(resp)
	return true, nil
}
