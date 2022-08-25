package providers

import "fmt"

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

func (e *Telegram) SendTele(phone string, content string) {
	fmt.Println("phone:\t", phone)
	fmt.Println("content:\t", content)
}
