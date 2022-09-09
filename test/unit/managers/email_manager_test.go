package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/cmd/managers"
	t "auth-plus-notification/test/unit/mocks"
)

type EmailManagerTestSuite struct {
	suite.Suite
}

func (suite *EmailManagerTestSuite) Test_succeed_when_choosing_sendgrid() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)
	const number = 0.1
	emailManager := m.NewRandomEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, sendgridMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *EmailManagerTestSuite) Test_succeed_when_choosing_onesignal() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)

	emailManager := m.NewRandomEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	const number = 0.4
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, onesignalMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *EmailManagerTestSuite) Test_succeed_when_choosing_mailgun() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)

	emailManager := m.NewRandomEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	const number = 0.7
	provider, err := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, mailgunMocked)
	assert.Equal(suite.T(), err, nil)
}

func TestEmailManager(t *testing.T) {
	suite.Run(t, new(EmailManagerTestSuite))
}
