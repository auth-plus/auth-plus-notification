package driven

type SendingWhatsapp interface {
	SendWhats(phone string, content string) (bool, error)
}
