package driven

type SendingWhatsapp interface {
	Send(phone string, content string)
}
