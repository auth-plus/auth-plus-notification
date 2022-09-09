package driven

// SendingPushNotification is a interface that must abstract how provider can send
type SendingPushNotification interface {
	SendPN(deviceID string, title string, content string) (bool, error)
}
