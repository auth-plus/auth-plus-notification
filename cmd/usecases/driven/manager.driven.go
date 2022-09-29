// Package driven contains all interfaces that usecases use
package driven

// Manager is a interface that must abstract whats input should be used for choosing a provider
type Manager[Provider any, Input any] interface {
	GetInput() (Input, error)
	ChooseProvider(number Input) (Provider, error)
}
