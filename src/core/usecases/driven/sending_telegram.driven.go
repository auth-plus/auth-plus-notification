package driven

type SendingTelegram interface {
	SendTele(phone string, content string)
}
