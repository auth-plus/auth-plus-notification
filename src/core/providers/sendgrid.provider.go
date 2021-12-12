package providers

import "fmt"

type Sendgrid struct {
	url   string
	token string
}

func NewSendgrid() *Sendgrid {
	instance := new(Sendgrid)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Sendgrid) Send(email string, content string) {
	fmt.Println("email Id:\t", email)
	fmt.Println("content Id:\t", content)
}
