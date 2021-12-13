package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type WhatsappUsecase struct {
	sendingWhatsapp d.SendingWhatsapp
}

func NewWhatsappUsecase() *WhatsappUsecase {
	instance := new(WhatsappUsecase)
	return instance
}

func (e *WhatsappUsecase) Send(phone string, content string) {
	e.sendingWhatsapp.Send(phone, content)
}
