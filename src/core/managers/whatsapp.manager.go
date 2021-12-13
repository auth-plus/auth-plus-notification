package managers

import (
	d "auth-plus-notification/core/usecases/driven"
)

//ENUM for Providers
type EnumWhatsappProvider string

const (
	Whatsapp EnumWhatsappProvider = "Whatsapp"
)

//Class for WhatsappManager
type WhatsappManager struct {
	whatsapp d.SendingWhatsapp
}

func NewWhatsappManager(whatsapp d.SendingWhatsapp) *WhatsappManager {
	return &WhatsappManager{whatsapp: whatsapp}
}

func (e *WhatsappManager) Send(phone string, content string) {
	choosedProvider := chooseWhatsappProvider(phone, content)
	switch choosedProvider {
	case "SendGrid":
		e.whatsapp.Send(phone, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseWhatsappProvider(phone string, content string) EnumWhatsappProvider {
	return "Whatsapp"
}
