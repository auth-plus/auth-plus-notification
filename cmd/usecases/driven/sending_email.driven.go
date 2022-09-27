package driven

// SendingEmail is a interface that must abstract how provider can send
type SendingEmail interface {
	SendEmail(email string, content string) error
}
