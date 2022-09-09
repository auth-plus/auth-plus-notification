package driven

// SendingWhatsapp is a interface that must abstract how provider can send
type SendingWhatsapp interface {
	SendWhats(phone string, content string) (bool, error)
}
