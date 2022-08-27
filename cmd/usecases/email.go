package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type EmailUsecase struct {
	sendingEmail d.SendingEmail
}

func NewEmailUsecase(sendingEmail d.SendingEmail) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.sendingEmail = sendingEmail
	return instance
}

func (e *EmailUsecase) Send(email string, content string) (bool, error) {
	return e.sendingEmail.SendEmail(email, content)
}
