// Package driven contains all interfaces that usecases use
package driven

// EmailManager is a interface that must abstract whats input should be used for choosing a provider
type EmailManager interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (SendingEmail, error)
}
