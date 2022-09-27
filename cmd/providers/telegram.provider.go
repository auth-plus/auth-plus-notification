package providers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Telegram struct must contains all private property to work
type Telegram struct {
	url   string
	token string
}

// NewTelegram for instanciate a telegram provider
func NewTelegram() *Telegram {
	instance := new(Telegram)
	instance.url = ""
	instance.token = ""
	return instance
}

// SendTele implementation of SendingTelegram
func (e *Telegram) SendTele(chatID int64, text string) error {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		return err
	}
	return nil

}
