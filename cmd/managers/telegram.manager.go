package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type TelegramManager struct {
	telegram d.SendingTelegram
}

func NewTelegramManager(telegram d.SendingTelegram) *TelegramManager {
	return &TelegramManager{telegram: telegram}
}

func (e *TelegramManager) SendTele(chatId int64, text string) {
	provider := e.chooseTelegramProvider(chatId, text)
	provider.SendTele(chatId, text)
}

// Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *TelegramManager) chooseTelegramProvider(chatId int64, text string) d.SendingTelegram {
	return e.telegram
}
