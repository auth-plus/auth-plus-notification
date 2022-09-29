package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/cmd/managers"
	t "auth-plus-notification/test/mocks"
)

type RandomSmsManagerTestSuite struct {
	suite.Suite
}

func (suite *RandomSmsManagerTestSuite) Test_succeed_when_choosing_sns() {
	snsMocked := new(t.SnsMocked)
	onesignalMocked := new(t.OnesignalMocked)
	const number = 0.7
	smsManager := m.NewRandomSmsManager(snsMocked, onesignalMocked)
	provider, err := smsManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, snsMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *RandomSmsManagerTestSuite) Test_succeed_when_choosing_onesignal() {
	snsMocked := new(t.SnsMocked)
	onesignalMocked := new(t.OnesignalMocked)

	smsManager := m.NewRandomSmsManager(snsMocked, onesignalMocked)
	const number = 0.1
	provider, err := smsManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, onesignalMocked)
	assert.Equal(suite.T(), err, nil)
}

func TestSmsManager(t *testing.T) {
	suite.Run(t, new(RandomPushNotificationManagerTestSuite))
}
