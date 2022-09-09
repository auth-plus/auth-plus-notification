package mock

import (
	d "auth-plus-notification/cmd/usecases/driven"

	"github.com/stretchr/testify/mock"
)

type RandomEmailManagerMocked struct {
	mock.Mock
}

func (m *RandomEmailManagerMocked) GetInput() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}
func (m *RandomEmailManagerMocked) ChooseProvider(number float64) (d.SendingEmail, error) {
	args := m.Called(number)
	return args.Get(0).(*SendgridMocked), args.Error(1)
}

type RandomPushNotificationManagerMocked struct {
	mock.Mock
}

func (m *RandomPushNotificationManagerMocked) GetInput() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}

func (m *RandomPushNotificationManagerMocked) ChooseProvider(number float64) (d.SendingPushNotification, error) {
	args := m.Called(number)
	return args.Get(0).(*FirebaseMocked), args.Error(1)
}
