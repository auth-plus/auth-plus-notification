package managers

import (
	se "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
)

type EmailManager struct {
	sendgrid  se.SendingEmail
	mailgun   se.SendingEmail
	onesignal se.SendingEmail
}

func NewEmailManager(sendgrid se.SendingEmail, mailgun se.SendingEmail, onesignal se.SendingEmail) *EmailManager {
	return &EmailManager{sendgrid: sendgrid, mailgun: mailgun, onesignal: onesignal}
}

func (e *EmailManager) SendEmail(email string, content string) (bool, error) {
	provider := e.chooseEmailManager(email, content)
	return provider.SendEmail(email, content)
}

// Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *EmailManager) chooseEmailManager(email string, content string) se.SendingEmail {
	random := rand.Float64()
	if random < 0.333 {
		return e.sendgrid
	}
	if random < 0.666 {
		return e.onesignal
	}
	return e.mailgun
}
