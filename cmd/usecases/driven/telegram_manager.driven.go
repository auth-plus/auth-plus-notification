package driven

// TelegramManager is a interface that must abstract whats input should be used for choosing a provider
type TelegramManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingTelegram, error)
}
