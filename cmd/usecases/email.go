package usecases

import (
	se "auth-plus-notification/cmd/usecases/driven"
)

type EmailUsecase struct {
	manager se.EmailManager
}

func NewEmailUsecase(manager se.EmailManager) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.manager = manager
	return instance
}

func (e *EmailUsecase) Send(email string, content string) (bool, error) {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return false, errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return false, errC
	}
	return provider.SendEmail(email, content)
}
