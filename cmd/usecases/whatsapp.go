package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type WhatsappUsecase struct {
	manager d.WhatsappManager
}

func NewWhatsappUsecase(manager d.WhatsappManager) *WhatsappUsecase {
	instance := new(WhatsappUsecase)
	instance.manager = manager
	return instance
}

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
