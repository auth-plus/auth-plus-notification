package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/internal/managers"
	t "auth-plus-notification/test/mocks"
)

type IPWarmimgEmailManagerTestSuite struct {
	suite.Suite
}

func (suite *IPWarmimgEmailManagerTestSuite) Test_succeed_when_choosing_sendgrid() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)
	input := m.IPWarmingInput{Sendgrid: 50, Mailgun: 50, Onesignal: 75}
	emailManager := m.NewIPWarmimgmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	provider, err := emailManager.ChooseProvider(input)
	assert.Equal(suite.T(), provider, sendgridMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *IPWarmimgEmailManagerTestSuite) Test_succeed_when_choosing_onesignal() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)
	input := m.IPWarmingInput{Sendgrid: 100, Mailgun: 50, Onesignal: 20}
	emailManager := m.NewIPWarmimgmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	provider, err := emailManager.ChooseProvider(input)
	assert.Equal(suite.T(), provider, onesignalMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *IPWarmimgEmailManagerTestSuite) Test_succeed_when_choosing_mailgun() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)

	input := m.IPWarmingInput{Sendgrid: 100, Mailgun: 50, Onesignal: 75}
	emailManager := m.NewIPWarmimgmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	provider, err := emailManager.ChooseProvider(input)
	assert.Equal(suite.T(), provider, mailgunMocked)
	assert.Equal(suite.T(), err, nil)
}

func (suite *IPWarmimgEmailManagerTestSuite) Test_succeed_when_getting_input() {
	sendgridMocked := new(t.SendgridMocked)
	mailgunMocked := new(t.MailgunMocked)
	onesignalMocked := new(t.OnesignalMocked)
	emailManager := m.NewIPWarmimgmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	input, err := emailManager.GetInput()
	assert.Equal(suite.T(), input.Sendgrid, 100)
	assert.Equal(suite.T(), input.Mailgun, 50)
	assert.Equal(suite.T(), input.Onesignal, 75)
	assert.Equal(suite.T(), err, nil)
}

func TestIPWarmimgEmailManager(t *testing.T) {
	suite.Run(t, new(IPWarmimgEmailManagerTestSuite))
}
