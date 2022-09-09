package driven

type SmsManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingSms, error)
}
