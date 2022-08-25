package providers

import "fmt"

type Whatsapp struct {
	url   string
	token string
}

func NewWhatsapp() *Whatsapp {
	instance := new(Whatsapp)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Whatsapp) SendWhats(phone string, content string) {
	fmt.Println("phone:\t", phone)
	fmt.Println("content:\t", content)
}
