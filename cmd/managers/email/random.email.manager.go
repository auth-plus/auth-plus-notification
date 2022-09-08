package managers

import (
	se "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
	"time"
)

type RandomEmailManager struct {
	sendgrid  se.SendingEmail
	mailgun   se.SendingEmail
	onesignal se.SendingEmail
}

func NewRandomEmailManager(sendgrid se.SendingEmail, mailgun se.SendingEmail, onesignal se.SendingEmail) *RandomEmailManager {
	instance := new(RandomEmailManager)
	instance.mailgun = mailgun
	instance.onesignal = onesignal
	instance.sendgrid = sendgrid
	return instance
}

// Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *RandomEmailManager) ChooseProvider(number float64) (se.SendingEmail, error) {
	if number < 0.333 {
		return e.sendgrid, nil
	}
	if number < 0.666 {
		return e.onesignal, nil
	}
	return e.mailgun, nil
}

func (e *RandomEmailManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
