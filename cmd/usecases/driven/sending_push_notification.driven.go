package driven

type SendingPushNotification interface {
	SendPN(deviceID string, title string, content string) (bool, error)
}
