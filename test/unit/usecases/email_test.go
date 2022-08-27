package test

import (
	"errors"
	"fmt"
	"testing"

	u "auth-plus-notification/cmd/usecases"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmailUsecaseTestSuite struct {
	suite.Suite
}

type EmailManagerMocked struct {
	mock.Mock
}

func (m *EmailManagerMocked) SendEmail(email string, content string) (bool, error) {
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

	emailManager := new(EmailManagerMocked)
	emailManager.On("SendEmail", mockData.Email, mockData.Sentence).Return(true, nil)

	emailUsecase := u.NewEmailUsecase(emailManager)
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

	mockedErr := errors.New("Provider timeout")
	emailManager := new(EmailManagerMocked)
	emailManager.On("SendEmail", mockData.Email, mockData.Sentence).Return(false, mockedErr)

	emailUsecase := u.NewEmailUsecase(emailManager)
	resp, err := emailUsecase.Send(mockData.Email, mockData.Sentence)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err, mockedErr)
}

func TestEmailUsecase(t *testing.T) {

	suite.Run(t, new(EmailUsecaseTestSuite))

}
