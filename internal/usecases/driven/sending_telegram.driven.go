package driven

import "context"

// SendingTelegram is a interface that must abstract how provider can send
type SendingTelegram interface {
	SendTele(ctx context.Context, chatID int64, text string) error
}
