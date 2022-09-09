package providers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	url   string
	token string
}

func NewTelegram() *Telegram {
	instance := new(Telegram)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Telegram) SendTele(chatId int64, text string) (bool, error) {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		return false, err
	}
	return true, nil

}
