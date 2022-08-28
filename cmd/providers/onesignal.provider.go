package providers

import "fmt"

type OneSignal struct {
	url   string
	token string
}

func NewOneSignal() *OneSignal {
	instance := new(OneSignal)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *OneSignal) SendEmail(email string, content string) (bool, error) {
	fmt.Println("email Id:\t", email)
	fmt.Println("content Id:\t", content)
	return true, nil
}

func (e *OneSignal) SendPN(deviceId string, title string, content string) {
	fmt.Println("device Id:\t", deviceId)
	fmt.Println("title:\t", title)
	fmt.Println("content:\t", content)
}

func (e *OneSignal) SendSms(phone string, content string) {
	fmt.Println("phone:\t", phone)
	fmt.Println("content:\t", content)
}
