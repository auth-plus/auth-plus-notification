package driven

type SendingTelegram interface {
	Send(phone string, content string)
}
