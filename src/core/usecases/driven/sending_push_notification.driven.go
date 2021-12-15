package driven

type SendingPushNotification interface {
	SendPN(deviceId string, title string, content string)
}
