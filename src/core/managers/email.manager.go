package managers

import (
	se "auth-plus-notification/core/usecases/driven"
	"math/rand"
)

//Class for EmailManager
type EmailManager struct {
	sendgrid  se.SendingEmail
	mailgun   se.SendingEmail
	onesignal se.SendingEmail
}

func NewEmailManager(sendgrid se.SendingEmail, mailgun se.SendingEmail, onesignal se.SendingEmail) *EmailManager {
	return &EmailManager{sendgrid: sendgrid, mailgun: mailgun, onesignal: onesignal}
}

func (e *EmailManager) SendEmail(email string, content string) {
	choosedProvider := chooseEmailManager(email, content)
	switch choosedProvider {
	case "SendGrid":
		e.sendgrid.SendEmail(email, content)
	case "Mailgun":
		e.mailgun.SendEmail(email, content)
	case "OneSignal":
		e.onesignal.SendEmail(email, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseEmailManager(email string, content string) EnumEmailProvider {
	random := rand.Float64()
	if random < 0.333 {
		return "SendGrid"
	}
	if random < 0.666 {
		return "OneSignal"
	}
	return "Mailgun"
}
