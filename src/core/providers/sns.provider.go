package providers

import "fmt"

type SNS struct {
	url   string
	token string
}

func NewSNS() *SNS {
	instance := new(SNS)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *SNS) SendSms(phone string, content string) {
	fmt.Println("phone:\t", phone)
	fmt.Println("content:\t", content)
}
