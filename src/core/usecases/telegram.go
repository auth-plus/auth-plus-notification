package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type TelegramUsecase struct {
	sendingTelegram d.SendingTelegram
}

func NewTelegramUsecase() *TelegramUsecase {
	instance := new(TelegramUsecase)
	return instance
}

func (e *TelegramUsecase) Send(phone string, content string) {
	e.sendingTelegram.Send(phone, content)
}
