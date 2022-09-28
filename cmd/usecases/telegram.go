package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

// TelegramUsecase dependencies
type TelegramUsecase struct {
	manager d.TelegramManager
}

// NewTelegramUsecase for instanciate a Telegram usecase
func NewTelegramUsecase(manager d.TelegramManager) *TelegramUsecase {
	instance := new(TelegramUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an telegram message by using manager on dependecy
func (e *TelegramUsecase) Send(chatID int64, text string) error {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return errC
	}
	return provider.SendTele(chatID, text)
}
