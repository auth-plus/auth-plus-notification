package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
	"time"
)

type randomEmailManager struct {
	sendgrid  d.SendingEmail
	mailgun   d.SendingEmail
	onesignal d.SendingEmail
}

// NewRandomEmailManager is a function for intanciate a pointer for EmailManager
func NewRandomEmailManager(sendgrid d.SendingEmail, mailgun d.SendingEmail, onesignal d.SendingEmail) *randomEmailManager {
	instance := new(randomEmailManager)
	instance.mailgun = mailgun
	instance.onesignal = onesignal
	instance.sendgrid = sendgrid
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *randomEmailManager) ChooseProvider(number float64) (d.SendingEmail, error) {
	if number < 0.333 {
		return e.sendgrid, nil
	}
	if number < 0.666 {
		return e.onesignal, nil
	}
	return e.mailgun, nil
}

// GetInput is a function that generate a random number
func (e *randomEmailManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
