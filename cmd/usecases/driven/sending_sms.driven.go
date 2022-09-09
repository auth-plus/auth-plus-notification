package driven

// SendingSms is a interface that must abstract how provider can send
type SendingSms interface {
	SendSms(phone string, content string) (bool, error)
}
