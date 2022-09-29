// Package driven contains all interfaces that usecases use
package driven

// Manager is a interface that must abstract whats input should be used for choosing a provider
type Manager[P any] interface {
	GetInput() (float64, error)
	ChooseProvider(number float64) (P, error)
}
