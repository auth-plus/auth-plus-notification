package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type SmsUsecase struct {
	manager d.SmsManager
}

func NewSmsUsecase(manager d.SmsManager) *SmsUsecase {
	instance := new(SmsUsecase)
	instance.manager = manager
	return instance
}

func (e *SmsUsecase) Send(phone string, content string) (bool, error) {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return false, errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return false, errC
	}
	return provider.SendSms(phone, content)
}
