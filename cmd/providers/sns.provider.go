package providers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SNS struct must contains all private property to work
type SNS struct {
	url   string
	token string
}

// NewSNS for instanciate a sns provider
func NewSNS() *SNS {
	instance := new(SNS)
	instance.url = ""
	instance.token = ""
	return instance
}

// SendSms implementation of SendingSms
func (e *SNS) SendSms(phone string, content string) error {
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
		return err
	}
	fmt.Println(resp)
	return nil
}
