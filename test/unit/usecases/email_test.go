package test

import (
	u "auth-plus-notification/cmd/usecases"
	se "auth-plus-notification/cmd/usecases/driven"
	"errors"
	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmailUsecaseTestSuite struct {
	suite.Suite
}

type RandomEmailManagerMocked struct {
	mock.Mock
}
type SendgridMocked struct {
	mock.Mock
}

func (m *RandomEmailManagerMocked) GetInput() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}
func (m *RandomEmailManagerMocked) ChooseProvider(number float64) (se.SendingEmail, error) {
	args := m.Called(number)
	return args.Get(0).(*SendgridMocked), args.Error(1)
}

func (m *SendgridMocked) SendEmail(email string, content string) (bool, error) {
	args := m.Called(email, content)
	return args.Bool(0), args.Error(1)
}

type MockedData struct {
	Email    string `faker:"email"`
	Sentence string `faker:"sentence"`
}

func (suite *EmailUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	sendgridMocked := new(SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(true, nil)

	var number float64 = 0.4
	randomEmailManager := new(RandomEmailManagerMocked)
	randomEmailManager.On("GetInput").Return(number, nil)
	randomEmailManager.On("ChooseProvider", number).Return(sendgridMocked, nil)

	emailUsecase := u.NewEmailUsecase(randomEmailManager)
	resp, err := emailUsecase.Send(mockData.Email, mockData.Sentence)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func (suite *EmailUsecaseTestSuite) Test_fail_when_sending() {
	mockData := MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	sendgridMocked := new(SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Sentence).Return(false, errors.New("failed"))

	var number float64 = 0.4
	randomEmailManager := new(RandomEmailManagerMocked)
	randomEmailManager.On("GetInput").Return(number, nil)
	randomEmailManager.On("ChooseProvider", number).Return(sendgridMocked, nil)

	emailUsecase := u.NewEmailUsecase(randomEmailManager)
	resp, err := emailUsecase.Send(mockData.Email, mockData.Sentence)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestEmailUsecase(t *testing.T) {

	suite.Run(t, new(EmailUsecaseTestSuite))

}
