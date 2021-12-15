package providers

import "fmt"

type Mailgun struct {
	url   string
	token string
}

func NewMailgun() *Mailgun {
	instance := new(Mailgun)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Mailgun) SendEmail(email string, content string) {
	fmt.Println("email Id:\t", email)
	fmt.Println("content Id:\t", content)
}
