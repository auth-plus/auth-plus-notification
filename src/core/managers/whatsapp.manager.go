package managers

import (
	d "auth-plus-notification/core/usecases/driven"
)

//Class for WhatsappManager
type WhatsappManager struct {
	whatsapp d.SendingWhatsapp
}

func NewWhatsappManager(whatsapp d.SendingWhatsapp) *WhatsappManager {
	return &WhatsappManager{whatsapp: whatsapp}
}

func (e *WhatsappManager) SendWhats(phone string, content string) {
	provider := e.chooseWhatsappProvider(phone, content)
	provider.SendWhats(phone, content)
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *WhatsappManager) chooseWhatsappProvider(phone string, content string) d.SendingWhatsapp {
	return e.whatsapp
}
