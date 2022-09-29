package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

// WhatsappUsecase dependencies
type WhatsappUsecase struct {
	manager d.Manager[d.SendingWhatsapp, float64]
}

// NewWhatsappUsecase for instanciate a whatsapp usecase
func NewWhatsappUsecase(manager d.Manager[d.SendingWhatsapp, float64]) *WhatsappUsecase {
	instance := new(WhatsappUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an whatsapp by using manager on dependecy
func (e *WhatsappUsecase) Send(phone string, content string) error {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return errC
	}
	return provider.SendWhats(phone, content)
}
