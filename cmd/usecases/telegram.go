package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type TelegramUsecase struct {
	sendingTelegram d.SendingTelegram
}

func NewTelegramUsecase(sendingTelegram d.SendingTelegram) *TelegramUsecase {
	instance := new(TelegramUsecase)
	instance.sendingTelegram = sendingTelegram
	return instance
}

func (e *TelegramUsecase) Send(chatId int64, text string) {
	e.sendingTelegram.SendTele(chatId, text)
}
