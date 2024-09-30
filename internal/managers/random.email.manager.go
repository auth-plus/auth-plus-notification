// Package managers contains all implementations of managers
package managers

import (
	d "auth-plus-notification/internal/usecases/driven"
	"math/rand"
)

// RandomEmailManager must contains all provider that could be choosen
type RandomEmailManager struct {
	sendgrid  d.SendingEmail
	mailgun   d.SendingEmail
	onesignal d.SendingEmail
}

// NewRandomEmailManager is a function for intanciate a pointer for EmailManager
func NewRandomEmailManager(sendgrid d.SendingEmail, mailgun d.SendingEmail, onesignal d.SendingEmail) *RandomEmailManager {
	instance := new(RandomEmailManager)
	instance.mailgun = mailgun
	instance.onesignal = onesignal
	instance.sendgrid = sendgrid
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *RandomEmailManager) ChooseProvider(number float64) (d.SendingEmail, error) {
	if number < 0.333 {
		return e.sendgrid, nil
	}
	if number < 0.666 {
		return e.onesignal, nil
	}
	return e.mailgun, nil
}

// GetInput is a function that generate a random number
func (e *RandomEmailManager) GetInput() (float64, error) {
	return rand.Float64(), nil
}
