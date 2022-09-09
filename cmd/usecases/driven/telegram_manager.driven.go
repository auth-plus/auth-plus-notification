package driven

type TelegramManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingTelegram, error)
}
