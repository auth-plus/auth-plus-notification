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

func (e *EmailManager) SendEmail(email string, content string) {
	choosedProvider := chooseEmailManager(email, content)
	switch choosedProvider {
	case "SendGrid":
		e.sendgrid.SendEmail(email, content)
	case "Mailgun":
		e.mailgun.SendEmail(email, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseEmailManager(email string, content string) EnumEmailProvider {
	if rand.Float64() > 0.5 {
		return "SendGrid"
	}
	return "Mailgun"
}
