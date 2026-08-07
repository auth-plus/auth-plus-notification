package mock

import (
	"context"

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
func (m *FirebaseMocked) SendPN(ctx context.Context, deviceID string, title string, content string) error {
	args := m.Called(ctx, deviceID, title, content)
	return args.Error(0)
}

// SendEmail mocked method for MailgunMocked
func (m *MailgunMocked) SendEmail(ctx context.Context, email string, _ string, content string) error {
	args := m.Called(ctx, email, content)
	return args.Error(0)
}

// SendEmail mocked method for OnesignalMocked
func (m *OnesignalMocked) SendEmail(ctx context.Context, email string, _ string, content string) error {
	args := m.Called(ctx, email, content)
	return args.Error(0)
}

// SendPN mocked method for OnesignalMocked
func (m *OnesignalMocked) SendPN(ctx context.Context, deviceID string, title string, content string) error {
	args := m.Called(ctx, deviceID, title, content)
	return args.Error(0)
}

// SendSms mocked method for OnesignalMocked
func (m *OnesignalMocked) SendSms(ctx context.Context, phone string, content string) error {
	args := m.Called(ctx, phone, content)
	return args.Error(0)
}

// SendEmail mocked method for SendgridMocked
func (m *SendgridMocked) SendEmail(ctx context.Context, email string, _ string, content string) error {
	args := m.Called(ctx, email, content)
	return args.Error(0)
}

// SendSms mocked method for SnsMocked
func (m *SnsMocked) SendSms(ctx context.Context, phone string, content string) error {
	args := m.Called(ctx, phone, content)
	return args.Error(0)
}

// SendTele mocked method for TelegramMocked
func (m *TelegramMocked) SendTele(ctx context.Context, chatID int64, text string) error {
	args := m.Called(ctx, chatID, text)
	return args.Error(0)
}

// SendWhats mocked method for TwilioMocked
func (m *TwilioMocked) SendWhats(ctx context.Context, phone string, content string) error {
	args := m.Called(ctx, phone, content)
	return args.Error(0)
}
