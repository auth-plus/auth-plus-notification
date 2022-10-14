package providers

import (
	"errors"
	"fmt"
	"log"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Twilio struct must contains all private property to work
type Twilio struct {
	url   string
	token string
}

// NewTwilio for instanciate a Twilio provider
func NewTwilio() *Twilio {
	instance := new(Twilio)
	instance.url = ""
	instance.token = ""
	return instance
}

// SendWhats implementation of SendingWhatsapp (https://www.twilio.com/blog/send-whatsapp-message-30-seconds-golang)
func (e *Twilio) SendWhats(phone string, content string) error {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:%s", phone))
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody(content)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println("TwilioError:", err)
		return errors.New("TwilioProvider: something went wrong")
	}
	return nil
}
