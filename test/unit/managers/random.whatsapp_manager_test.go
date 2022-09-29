package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/cmd/managers"
	t "auth-plus-notification/test/mocks"
)

type RandomWhatsappManagerTestSuite struct {
	suite.Suite
}

func (suite *RandomWhatsappManagerTestSuite) Test_succeed_when_choosing_twilio() {
	twilioMocked := new(t.TwilioMocked)
	const number = 0.7
	smsManager := m.NewRandomWhatsappManager(twilioMocked)
	provider, err := smsManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, twilioMocked)
	assert.Equal(suite.T(), err, nil)
}

func TestWhatsappManager(t *testing.T) {
	suite.Run(t, new(RandomPushNotificationManagerTestSuite))
}
