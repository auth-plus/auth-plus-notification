package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

// SmsUsecase dependencies
type SmsUsecase struct {
	manager d.SmsManager
}

// NewSmsUsecase for instanciate a sms usecase
func NewSmsUsecase(manager d.SmsManager) *SmsUsecase {
	instance := new(SmsUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an sms by using manager on dependecy
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
