package providers

import "fmt"

type Firebase struct {
	url   string
	token string
}

func NewFirebase() *Firebase {
	instance := new(Firebase)
	instance.url = ""
	instance.token = ""
	return instance
}

func (e *Firebase) SendPN(deviceId string, title string, content string) {
	fmt.Println("deviceId:\t", deviceId)
	fmt.Println("title:\t", title)
	fmt.Println("content:\t", content)
}
