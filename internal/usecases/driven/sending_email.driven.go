package driven

import "context"

// SendingEmail is a interface that must abstract how provider can send
type SendingEmail interface {
	SendEmail(ctx context.Context, email string, subject string, content string) error
}
