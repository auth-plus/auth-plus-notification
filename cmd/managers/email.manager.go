package managers

import (
	se "auth-plus-notification/cmd/usecases/driven"
)

type EmailManager struct {
	sendgrid  se.SendingEmail
	mailgun   se.SendingEmail
	onesignal se.SendingEmail
}

func NewEmailManager(sendgrid se.SendingEmail, mailgun se.SendingEmail, onesignal se.SendingEmail) *EmailManager {
	instance := new(EmailManager)
	instance.mailgun = mailgun
	instance.onesignal = onesignal
	instance.sendgrid = sendgrid
	return instance
}

// Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *EmailManager) ChooseProvider(random float64) se.SendingEmail {
	if random < 0.333 {
		return e.sendgrid
	}
	if random < 0.666 {
		return e.onesignal
	}
	return e.mailgun
}
