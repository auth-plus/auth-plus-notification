package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type SmsUsecase struct {
	sendingSms d.SendingSms
}

func NewSmsUsecase() *SmsUsecase {
	instance := new(SmsUsecase)
	return instance
}

func (e *SmsUsecase) Send(phone string, content string) {
	e.sendingSms.Send(phone, content)
}
