package managers

import (
	d "auth-plus-notification/core/usecases/driven"
)

//ENUM for Providers
type EnumTelegramProvider string

const (
	Telegram EnumTelegramProvider = "Telegram"
)

//Class for TelegramManager
type TelegramManager struct {
	telegram d.SendingTelegram
}

func NewTelegramManager(telegram d.SendingTelegram) *TelegramManager {
	return &TelegramManager{telegram: telegram}
}

func (e *TelegramManager) Send(phone string, content string) {
	choosedProvider := chooseTelegramProvider(phone, content)
	switch choosedProvider {
	case "SendGrid":
		e.telegram.Send(phone, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func chooseTelegramProvider(phone string, content string) EnumTelegramProvider {
	return "Telegram"
}
