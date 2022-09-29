package mock

import (
	"github.com/stretchr/testify/mock"
)

// ManagerMocked is a mock for Manager interface
type ManagerMocked[Provider any, Input any] struct {
	mock.Mock
}

// GetInput mocking of EmailManager method
func (m *ManagerMocked[Provider, Input]) GetInput() (Input, error) {
	args := m.Called()
	return args.Get(0).(Input), args.Error(1)
}

// ChooseProvider mocking of EmailManager method
func (m *ManagerMocked[Provider, Input]) ChooseProvider(ipt Input) (Provider, error) {
	args := m.Called(ipt)
	return args.Get(0).(Provider), args.Error(1)
}
