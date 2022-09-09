package driven

// SendingTelegram is a interface that must abstract how provider can send
type SendingTelegram interface {
	SendTele(chatID int64, text string) (bool, error)
}
