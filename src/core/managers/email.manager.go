package managers

import (
	se "auth-plus-notification/core/usecases/driven"
	"math/rand"
)

//Class for EmailManager
type EmailManager struct {
	sendgrid se.SendingEmail
	mailgun  se.SendingEmail
	braze    se.SendingEmail
}

func NewEmailManager(sendgrid se.SendingEmail, mailgun se.SendingEmail, braze se.SendingEmail) *EmailManager {
	return &EmailManager{sendgrid: sendgrid, mailgun: mailgun, braze: braze}
}

func (e *EmailManager) SendEmail(email string, content string) {
	choosedProvider := chooseEmailManager(email, content)
	switch choosedProvider {
	case "SendGrid":
		e.sendgrid.SendEmail(email, content)
	case "Mailgun":
		e.mailgun.SendEmail(email, content)
	case "Braze":
		e.braze.SendEmail(email, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseEmailManager(email string, content string) EnumEmailProvider {
	switch {
	case rand.Float64() < 0.33:
		return "SendGrid"

	case rand.Float64() < 0.66:
		return "Mailgun"

	default:
		return "Braze"
	}

}
