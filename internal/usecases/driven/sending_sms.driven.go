package driven

import "context"

// SendingSms is a interface that must abstract how provider can send
type SendingSms interface {
	SendSms(ctx context.Context, phone string, content string) error
}
