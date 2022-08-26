package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

//Class for TelegramManager
type TelegramManager struct {
	telegram d.SendingTelegram
}

func NewTelegramManager(telegram d.SendingTelegram) *TelegramManager {
	return &TelegramManager{telegram: telegram}
}

func (e *TelegramManager) SendTele(phone string, content string) {
	provider := e.chooseTelegramProvider(phone, content)
	provider.SendTele(phone, content)
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *TelegramManager) chooseTelegramProvider(phone string, content string) d.SendingTelegram {
	return e.telegram
}
