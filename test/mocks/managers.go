package mock

import (
	"github.com/stretchr/testify/mock"
)

// ManagerMocked is a mock for Manager interface
type ManagerMocked[Provider any] struct {
	mock.Mock
}

// GetInput mocking of EmailManager method
func (m *ManagerMocked[P]) GetInput() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}

// ChooseProvider mocking of EmailManager method
func (m *ManagerMocked[Provider]) ChooseProvider(number float64) (Provider, error) {
	args := m.Called(number)
	return args.Get(0).(Provider), args.Error(1)
}
