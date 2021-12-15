package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type TelegramUsecase struct {
	sendingTelegram d.SendingTelegram
}

func NewTelegramUsecase(sendingTelegram d.SendingTelegram) *TelegramUsecase {
	instance := new(TelegramUsecase)
	instance.sendingTelegram = sendingTelegram
	return instance
}

func (e *TelegramUsecase) Send(phone string, content string) {
	e.sendingTelegram.SendTele(phone, content)
}
