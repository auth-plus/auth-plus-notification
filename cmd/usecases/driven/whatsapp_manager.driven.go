package driven

type WhatsappManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingWhatsapp, error)
}
