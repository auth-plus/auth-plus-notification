package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/cmd/managers"
	t "auth-plus-notification/test/unit/mocks"
)

type PushNotificationManagerTestSuite struct {
	suite.Suite
}

func (suite *PushNotificationManagerTestSuite) Test_succeed_when_choosing_firebase() {
	firebaseMocked := new(t.FirebaseMocked)
	onesignalMocked := new(t.OnesignalMocked)
	const number = 0.7
	emailManager := m.NewRandomPushNotificationManager(firebaseMocked, onesignalMocked)
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, firebaseMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *PushNotificationManagerTestSuite) Test_succeed_when_choosing_onesignal() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)

	emailManager := m.NewRandomEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	const number = 0.1
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, onesignalMocked)
	assert.Equal(suite.T(), err, nil)
}

func TestPushNOtificationManager(t *testing.T) {
	suite.Run(t, new(PushNotificationManagerTestSuite))
}
