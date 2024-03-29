package usecases

import (
	d "auth-plus-notification/internal/usecases/driven"
)

// SmsUsecase dependencies
type SmsUsecase struct {
	manager d.Manager[d.SendingSms, float64]
}

// NewSmsUsecase for instanciate a sms usecase
func NewSmsUsecase(manager d.Manager[d.SendingSms, float64]) *SmsUsecase {
	instance := new(SmsUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an sms by using manager on dependecy
func (e *SmsUsecase) Send(phone string, content string) error {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return errC
	}
	return provider.SendSms(phone, content)
}
