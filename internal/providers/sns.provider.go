package providers

import (
	config "auth-plus-notification/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SNS struct must contains all private property to work
type SNS struct {
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
}

// NewSNS for instanciate a sns provider
func NewSNS() *SNS {
	env := config.GetEnv()
	instance := new(SNS)
	instance.accessKeyID = env.Providers.Amazon.AccessKeyID
	instance.secretAccessKey = env.Providers.Amazon.SecretAccessKey
	instance.sessionToken = env.Providers.Amazon.SessionToken
	return instance
}

// SendSms implementation of SendingSms
func (e *SNS) SendSms(phone string, content string) error {
	sess, errInit := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if errInit != nil {
		return errInit
	}
	svc := sns.New(sess)
	params := &sns.PublishInput{
		Message:     aws.String(content),
		PhoneNumber: aws.String(phone),
	}
	_, errPub := svc.Publish(params)
	if errPub != nil {
		return errPub
	}
	return nil
}
