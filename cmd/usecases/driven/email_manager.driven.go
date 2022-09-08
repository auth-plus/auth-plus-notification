package driven

type EmailManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingEmail, error)
}
