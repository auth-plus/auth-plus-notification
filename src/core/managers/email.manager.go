package managers

import (
	se "auth-plus-notification/core/usecases/driven"
	"math/rand"
)

//ENUM for Providers
type EnumEmailProvider string

const (
	SendGrid EnumEmailProvider = "SendGrid"
	Mailgun  EnumEmailProvider = "Mailgun"
)

//Class for EmailManager
type EmailManager struct {
	sendgrid se.SendingEmail
	mailgun  se.SendingEmail
}

func NewEmailManager(sendgrid se.SendingEmail, mailgun se.SendingEmail) *EmailManager {
	return &EmailManager{sendgrid: sendgrid, mailgun: mailgun}
}

func (e *EmailManager) Send(email string, content string) {
	choosedProvider := Choose(email, content)
	switch choosedProvider {
	case "SendGrid":
		e.sendgrid.Send(email, content)
	case "Mailgun":
		e.mailgun.Send(email, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func Choose(email string, content string) EnumEmailProvider {
	if rand.Float64() > 0.5 {
		return "SendGrid"
	}
	return "Mailgun"
}
