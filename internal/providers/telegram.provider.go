package providers

import (
	config "auth-plus-notification/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Telegram struct must contains all private property to work
type Telegram struct {
	token string
}

// NewTelegram for instanciate a telegram provider
func NewTelegram() *Telegram {
	env := config.GetEnv()
	instance := new(Telegram)
	instance.token = env.Providers.Telegram.APIKey
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
		return errR
	}
	return nil

}
