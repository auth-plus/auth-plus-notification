package driven

// WhatsappManager is a interface that must abstract whats input should be used for choosing a provider
type WhatsappManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingWhatsapp, error)
}
