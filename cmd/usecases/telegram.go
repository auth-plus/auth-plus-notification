package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type TelegramUsecase struct {
	manager d.TelegramManager
}

func NewTelegramUsecase(manager d.TelegramManager) *TelegramUsecase {
	instance := new(TelegramUsecase)
	instance.manager = manager
	return instance
}

func (e *TelegramUsecase) Send(chatId int64, text string) (bool, error) {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return false, errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return false, errC
	}
	return provider.SendTele(chatId, text)
}
