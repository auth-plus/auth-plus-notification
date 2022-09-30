package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/cmd/managers"
	t "auth-plus-notification/test/mocks"
)

type RandomPushNotificationManagerTestSuite struct {
	suite.Suite
}

func (suite *RandomPushNotificationManagerTestSuite) Test_succeed_when_choosing_firebase() {
	firebaseMocked := new(t.FirebaseMocked)
	onesignalMocked := new(t.OnesignalMocked)
	const number = 0.7
	emailManager := m.NewRandomPushNotificationManager(firebaseMocked, onesignalMocked)
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, firebaseMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *RandomPushNotificationManagerTestSuite) Test_succeed_when_choosing_onesignal() {
	firebaseMocked := new(t.FirebaseMocked)
	onesignalMocked := new(t.OnesignalMocked)

	emailManager := m.NewRandomPushNotificationManager(firebaseMocked, onesignalMocked)
	const number = 0.1
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, onesignalMocked)
	assert.Equal(suite.T(), err, nil)
}

func TestRandomPushNotificationManager(t *testing.T) {
	suite.Run(t, new(RandomPushNotificationManagerTestSuite))
}
