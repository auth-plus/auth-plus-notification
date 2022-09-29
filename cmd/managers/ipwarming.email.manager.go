// Package managers contains all implementations of managers
package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

// IPWarmimgmailManager must contains all provider that could be choosen
type IPWarmimgmailManager struct {
	sendgrid  d.SendingEmail
	mailgun   d.SendingEmail
	onesignal d.SendingEmail
}

// NewIPWarmimgmailManager is a function for intanciate a pointer for EmailManager
func NewIPWarmimgmailManager(sendgrid d.SendingEmail, mailgun d.SendingEmail, onesignal d.SendingEmail) *IPWarmimgmailManager {
	instance := new(IPWarmimgmailManager)
	instance.mailgun = mailgun
	instance.onesignal = onesignal
	instance.sendgrid = sendgrid
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *IPWarmimgmailManager) ChooseProvider(input IPWarmingInput) (d.SendingEmail, error) {
	if input.Sendgrid < 70 {
		return e.sendgrid, nil
	}
	if input.Onesignal < 30 {
		return e.onesignal, nil
	}
	return e.mailgun, nil
}

// IPWarmingInput is a struct representing what database count IP warmimg
type IPWarmingInput struct {
	Sendgrid  int
	Mailgun   int
	Onesignal int
}

// GetInput is a function that generate a random number
func (e *IPWarmimgmailManager) GetInput() (IPWarmingInput, error) {
	input := IPWarmingInput{
		Sendgrid:  100,
		Mailgun:   50,
		Onesignal: 75,
	}
	return input, nil
}
