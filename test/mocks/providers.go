package mock

import (
	"github.com/stretchr/testify/mock"
)

// FirebaseMocked mocking of Firebase provider
type FirebaseMocked struct {
	mock.Mock
}

// MailgunMocked mocking of Mailgun provider
type MailgunMocked struct {
	mock.Mock
}

// OnesignalMocked mocking of Onesignal provider
type OnesignalMocked struct {
	mock.Mock
}

// SendgridMocked mocking of Sendgrid provider
type SendgridMocked struct {
	mock.Mock
}

// SnsMocked mocking of Sns provider
type SnsMocked struct {
	mock.Mock
}

// TelegramMocked mocking of Telegram provider
type TelegramMocked struct {
	mock.Mock
}

// TwilioMocked mocking of Twilio provider
type TwilioMocked struct {
	mock.Mock
}

// SendPN mocked method for FirebaseMocked
func (m *FirebaseMocked) SendPN(deviceID string, title string, content string) (bool, error) {
	args := m.Called(deviceID, title, content)
	return args.Bool(0), args.Error(1)
}

// SendEmail mocked method for MailgunMocked
func (m *MailgunMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

// SendEmail mocked method for MailgunMocked
func (m *OnesignalMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

// SendPN mocked method for OnesignalMocked
func (m *OnesignalMocked) SendPN(deviceID string, title string, content string) (bool, error) {
	args := m.Called(deviceID, title, content)
	return args.Bool(0), args.Error(1)
}

// SendSms mocked method for OnesignalMocked
func (m *OnesignalMocked) SendSms(phone string, content string) (bool, error) {
	args := m.Called(phone, content)
	return args.Bool(0), args.Error(1)
}

// SendEmail mocked method for SendgridMocked
func (m *SendgridMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

// SendSms mocked method for SnsMocked
func (m *SnsMocked) SendSms(phone string, content string) (bool, error) {
	args := m.Called(phone, content)
	return args.Bool(0), args.Error(1)
}

// SendTele mocked method for TelegramMocked
func (m *TelegramMocked) SendTele(chatID int64, text string) (bool, error) {
	args := m.Called(chatID, text)
	return args.Bool(0), args.Error(1)
}

// SendWhats mocked method for TwilioMocked
func (m *TwilioMocked) SendWhats(phone string, content string) (bool, error) {
	args := m.Called(phone, content)
	return args.Bool(0), args.Error(1)
}
