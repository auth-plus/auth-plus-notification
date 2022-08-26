package providers

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Twilio struct {
	url   string
	token string
}

func NewTwilio() *Twilio {
	instance := new(Twilio)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Twilio) SendWhats(phone string, content string) {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:<YOUR-PHONE-NUMBER-HERE>")
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello from Golang!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
	fmt.Println("phone:\t", phone)
	fmt.Println("content:\t", content)
}
