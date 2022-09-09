package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

// WhatsappUsecase dependencies
type WhatsappUsecase struct {
	manager d.WhatsappManager
}

// NewWhatsappUsecase for instanciate a whatsapp usecase
func NewWhatsappUsecase(manager d.WhatsappManager) *WhatsappUsecase {
	instance := new(WhatsappUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an whatsapp by using manager on dependecy
func (e *WhatsappUsecase) Send(phone string, content string) (bool, error) {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return false, errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return false, errC
	}
	return provider.SendWhats(phone, content)
}
