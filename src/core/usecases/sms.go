package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type SmsUsecase struct {
	sendingSms d.SendingSms
}

func NewSmsUsecase(sendingSms d.SendingSms) *SmsUsecase {
	instance := new(SmsUsecase)
	instance.sendingSms = sendingSms
	return instance
}

func (e *SmsUsecase) Send(phone string, content string) {
	e.sendingSms.SendSms(phone, content)
}
