package test

import (
	"testing"

	m "auth-plus-notification/cmd/managers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmailManagerTestSuite struct {
	suite.Suite
}

type SendgridMocked struct {
	mock.Mock
}

type MailgunMocked struct {
	mock.Mock
}

type OnesignalMocked struct {
	mock.Mock
}

func (m *SendgridMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

func (m *MailgunMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}
func (m *OnesignalMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

func (suite *EmailManagerTestSuite) Test_succeed_when_choosing_sendgrid() {
	sendgridMocked := new(SendgridMocked)
	mailgunMocked := new(MailgunMocked)
	onesignalMocked := new(OnesignalMocked)

	emailManager := m.NewEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	var number float64 = 0.1
	provider := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, sendgridMocked)
}

func (suite *EmailManagerTestSuite) Test_succeed_when_choosing_onesignal() {
	sendgridMocked := new(SendgridMocked)
	mailgunMocked := new(MailgunMocked)
	onesignalMocked := new(OnesignalMocked)

	emailManager := m.NewEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	var number float64 = 0.4
	provider := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, onesignalMocked)
}

func (suite *EmailManagerTestSuite) Test_succeed_when_choosing_mailgun() {
	sendgridMocked := new(SendgridMocked)
	mailgunMocked := new(MailgunMocked)
	onesignalMocked := new(OnesignalMocked)

	emailManager := m.NewEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	var number float64 = 0.7
	provider := emailManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, mailgunMocked)
}

func TestEmailManager(t *testing.T) {
	suite.Run(t, new(EmailManagerTestSuite))
}
