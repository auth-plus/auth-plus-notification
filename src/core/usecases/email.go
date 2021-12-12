package usecases

import (
	se "auth-plus-notification/core/usecases/driven"
)

type EmailUsecase struct {
	sendingEmail se.SendingEmail
}

func NewEmailUsecase(sendingEmail se.SendingEmail) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.sendingEmail = sendingEmail
	return instance
}
