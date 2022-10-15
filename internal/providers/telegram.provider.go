package providers

import (
	config "auth-plus-notification/config"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// Telegram struct must contains all private property to work
type Telegram struct {
	token  string
	logger *zap.Logger
}

// NewTelegram for instanciate a telegram provider
func NewTelegram() *Telegram {
	env := config.GetEnv()
	instance := new(Telegram)
	instance.token = env.Providers.Telegram.APIKey
	instance.logger = config.GetLogger()
	return instance
}

// SendTele implementation of SendingTelegram
func (e *Telegram) SendTele(chatID int64, text string) error {
	bot, errInit := tgbotapi.NewBotAPI(e.token)
	if errInit != nil {
		return errInit
	}
	msg := tgbotapi.NewMessage(chatID, text)
	_, errR := bot.Send(msg)
	if errR != nil {
		e.logger.Error(errR.Error())
		return errors.New("TelegramProvider: something went wrong")
	}
	return nil

}
