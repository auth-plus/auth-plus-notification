package usecases

import (
	se "auth-plus-notification/cmd/usecases/driven"
)

type EmailUsecase struct {
	EmailManager se.EmailManager
}

func NewEmailUsecase(emailManager se.EmailManager) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.EmailManager = emailManager
	return instance
}

func (e *EmailUsecase) Send(email string, content string) (bool, error) {
	number, errI := e.EmailManager.GetInput()
	if errI != nil {
		return false, errI
	}
	provider, errC := e.EmailManager.ChooseProvider(number)
	if errC != nil {
		return false, errC
	}
	return provider.SendEmail(email, content)
}
