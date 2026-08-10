package driven

import "context"

// SendingWhatsapp is a interface that must abstract how provider can send
type SendingWhatsapp interface {
	SendWhats(ctx context.Context, phone string, content string) error
}
