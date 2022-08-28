package test

import (
	"errors"
	"fmt"
	"testing"

	m "auth-plus-notification/cmd/managers"

	"github.com/bxcodec/faker"
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

type MockedData struct {
	Email    string `faker:"email"`
	Sentence string `faker:"sentence"`
}

func (suite *EmailManagerTestSuite) Test_succeed_when_sending() {
	mockData := MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	sendgridMocked := new(SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(true, nil)
	mailgunMocked := new(MailgunMocked)
	mailgunMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(true, nil)
	onesignalMocked := new(OnesignalMocked)
	onesignalMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(true, nil)

	emailUsecase := m.NewEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	resp, err := emailUsecase.SendEmail(mockData.Email, mockData.Sentence)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func (suite *EmailManagerTestSuite) Test_fail_when_sending() {
	mockedErr := errors.New("Provider timeout")
	mockData := MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	sendgridMocked := new(SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(false, mockedErr)
	mailgunMocked := new(MailgunMocked)
	mailgunMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(false, mockedErr)
	onesignalMocked := new(OnesignalMocked)
	onesignalMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(false, mockedErr)

	emailUsecase := m.NewEmailManager(sendgridMocked, mailgunMocked, onesignalMocked)
	resp, err := emailUsecase.SendEmail(mockData.Email, mockData.Sentence)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err, mockedErr)
}

func TestEmailManager(t *testing.T) {

	suite.Run(t, new(EmailManagerTestSuite))

}
