package mock

import (
	d "auth-plus-notification/cmd/usecases/driven"

	"github.com/stretchr/testify/mock"
)

// RandomEmailManagerMocked is a mock for RandomEmailManager
type RandomEmailManagerMocked struct {
	mock.Mock
}

// GetInput mocking of EmailManager method
func (m *RandomEmailManagerMocked) GetInput() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}

// ChooseProvider mocking of EmailManager method
func (m *RandomEmailManagerMocked) ChooseProvider(number float64) (d.SendingEmail, error) {
	args := m.Called(number)
	return args.Get(0).(*SendgridMocked), args.Error(1)
}

// RandomPushNotificationManagerMocked is a mock for RandomPushNotificationManager
type RandomPushNotificationManagerMocked struct {
	mock.Mock
}

// GetInput mocking of PushNotificationManager method
func (m *RandomPushNotificationManagerMocked) GetInput() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}

// ChooseProvider mocking of PushNotificationManager method
func (m *RandomPushNotificationManagerMocked) ChooseProvider(number float64) (d.SendingPushNotification, error) {
	args := m.Called(number)
	return args.Get(0).(*FirebaseMocked), args.Error(1)
}
