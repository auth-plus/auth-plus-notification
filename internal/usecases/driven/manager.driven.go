// Package driven contains all interfaces that usecases use
package driven

import "context"

// Manager is a interface that must abstract whats input should be used for choosing a provider
type Manager[Provider any, Input any] interface {
	GetInput(ctx context.Context) (Input, error)
	ChooseProvider(ctx context.Context, number Input) (Provider, error)
}
