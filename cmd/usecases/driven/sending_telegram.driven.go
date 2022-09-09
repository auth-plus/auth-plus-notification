package driven

type SendingTelegram interface {
	SendTele(chatId int64, text string) (bool, error)
}
