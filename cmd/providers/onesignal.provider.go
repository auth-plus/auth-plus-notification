package providers

import "fmt"

// OneSignal struct must contains all private property to work
type OneSignal struct {
	url   string
	token string
}

// NewOneSignal for instanciate a onesignal provider
func NewOneSignal() *OneSignal {
	instance := new(OneSignal)
	instance.url = ""
	instance.token = ""
	return instance
}

// SendEmail implementation of SendingEmail
func (e *OneSignal) SendEmail(email string, content string) error {
	fmt.Println("email Id:\t", email)
	fmt.Println("content Id:\t", content)
	return nil
}

// SendPN implementation of SendingPushNotification
func (e *OneSignal) SendPN(deviceID string, title string, content string) error {
	fmt.Println("device Id:\t", deviceID)
	fmt.Println("title:\t", title)
	fmt.Println("content:\t", content)
	return nil
}

// SendSms implementation of SendingSms
func (e *OneSignal) SendSms(phone string, content string) error {
	fmt.Println("phone:\t", phone)
	fmt.Println("content:\t", content)
	return nil
}
