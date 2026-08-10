package driven

import "context"

// SendingPushNotification is a interface that must abstract how provider can send
type SendingPushNotification interface {
	SendPN(ctx context.Context, deviceID string, title string, content string) error
}
