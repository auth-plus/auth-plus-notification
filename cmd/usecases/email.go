// Package usecases contains all usecases
package usecases

import (
	m "auth-plus-notification/cmd/managers"
	se "auth-plus-notification/cmd/usecases/driven"
)

// EmailUsecase dependencies
type EmailUsecase struct {
	manager se.Manager[se.SendingEmail, m.IPWarmingInput]
}

// NewEmailUsecase for instanciate a email usecase
func NewEmailUsecase(manager se.Manager[se.SendingEmail, m.IPWarmingInput]) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an email by using manager on dependecy
func (e *EmailUsecase) Send(email string, content string) error {
	input, errI := e.manager.GetInput()
	if errI != nil {
		return errI
	}
	provider, errC := e.manager.ChooseProvider(input)
	if errC != nil {
		return errC
	}
	return provider.SendEmail(email, content)
}
