package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type WhatsappManager struct {
	twilio d.SendingWhatsapp
}

func NewWhatsappManager(twilio d.SendingWhatsapp) *WhatsappManager {
	return &WhatsappManager{twilio: twilio}
}

func (e *WhatsappManager) SendWhats(phone string, content string) {
	provider := e.chooseWhatsappProvider(phone, content)
	provider.SendWhats(phone, content)
}

func (e *WhatsappManager) chooseWhatsappProvider(phone string, content string) d.SendingWhatsapp {
	return e.twilio
}
