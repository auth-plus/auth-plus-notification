package mock

import (
	"github.com/stretchr/testify/mock"
)

type FirebaseMocked struct {
	mock.Mock
}
type MailgunMocked struct {
	mock.Mock
}

type OnesignalMocked struct {
	mock.Mock
}
type SendgridMocked struct {
	mock.Mock
}
type SnsMocked struct {
	mock.Mock
}
type TelegramMocked struct {
	mock.Mock
}
type TwilioMocked struct {
	mock.Mock
}

func (m *FirebaseMocked) SendPN(deviceId string, title string, content string) (bool, error) {
	args := m.Called(deviceId, title, content)
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
func (m *OnesignalMocked) SendPN(deviceId string, title string, content string) (bool, error) {
	args := m.Called(deviceId, title, content)
	return args.Bool(0), args.Error(1)
}
func (m *OnesignalMocked) SendSms(phone string, content string) (bool, error) {
	args := m.Called(phone, content)
	return args.Bool(0), args.Error(1)
}

func (m *SendgridMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

func (m *SnsMocked) SendSms(phone string, content string) (bool, error) {
	args := m.Called(phone, content)
	return args.Bool(0), args.Error(1)
}

func (m *TelegramMocked) SendTele(chatId int64, text string) (bool, error) {
	args := m.Called(chatId, text)
	return args.Bool(0), args.Error(1)
}

func (m *TwilioMocked) SendWhats(phone string, content string) (bool, error) {
	args := m.Called(phone, content)
	return args.Bool(0), args.Error(1)
}
