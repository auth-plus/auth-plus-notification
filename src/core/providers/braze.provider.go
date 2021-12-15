package providers

import "fmt"

type Braze struct {
	url   string
	token string
}

func NewBraze() *Braze {
	instance := new(Braze)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Braze) SendPN(deviceId string, title string, content string) {
	fmt.Println("deviceId:\t", deviceId)
	fmt.Println("title:\t", title)
	fmt.Println("content:\t", content)
}

func (e *Braze) SendEmail(email string, content string) {
	fmt.Println("email Id:\t", email)
	fmt.Println("content Id:\t", content)
}
