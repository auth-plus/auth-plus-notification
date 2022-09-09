package driven

// SmsManager is a interface that must abstract whats input should be used for choosing a provider
type SmsManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingSms, error)
}
