package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type WhatsappUsecase struct {
	sendingWhatsapp d.SendingWhatsapp
}

func NewWhatsappUsecase(sendingWhatsapp d.SendingWhatsapp) *WhatsappUsecase {
	instance := new(WhatsappUsecase)
	instance.sendingWhatsapp = sendingWhatsapp
	return instance
}

func (e *WhatsappUsecase) Send(phone string, content string) {
	e.sendingWhatsapp.SendWhats(phone, content)
}
