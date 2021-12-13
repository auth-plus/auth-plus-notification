package driven

type SendingPushNotification interface {
	Send(deviceId string, title string, content string)
}
