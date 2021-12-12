package managers

import (
	se "auth-plus-notification/core/usecases/driven"
)

//ENUM for Providers
type EnumEmailProvider string

const (
	SendGrid EnumEmailProvider = "SendGrid"
)

//Class for EmailManager
type EmailManager struct {
	sendgrid se.SendingEmail
}

func NewEmailManager(sendgrid se.SendingEmail) *EmailManager {
	return &EmailManager{sendgrid: sendgrid}
}

func (e *EmailManager) Send(email string, content string) {
	choosedProvider := Choose(email, content)
	switch choosedProvider {
	case "SendGrid":
		e.Send(email, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func Choose(email string, content string) EnumEmailProvider {
	return "SendGrid"
}
