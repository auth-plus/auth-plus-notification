package test

import (
	u "auth-plus-notification/cmd/usecases"
	t "auth-plus-notification/test/mocks"

	"errors"
	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EmailUsecaseTestSuite struct {
	suite.Suite
}

func (suite *EmailUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	sendgridMocked := new(t.SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Content).Return(true, nil)

	const number = 0.4
	randomEmailManager := new(t.RandomEmailManagerMocked)
	randomEmailManager.On("GetInput").Return(number, nil)
	randomEmailManager.On("ChooseProvider", number).Return(sendgridMocked, nil)

	emailUsecase := u.NewEmailUsecase(randomEmailManager)
	resp, err := emailUsecase.Send(mockData.Email, mockData.Content)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func (suite *EmailUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	sendgridMocked := new(t.SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Content).Return(false, errors.New("failed"))

	const number = 0.4
	randomEmailManager := new(t.RandomEmailManagerMocked)
	randomEmailManager.On("GetInput").Return(number, nil)
	randomEmailManager.On("ChooseProvider", number).Return(sendgridMocked, nil)

	emailUsecase := u.NewEmailUsecase(randomEmailManager)
	resp, err := emailUsecase.Send(mockData.Email, mockData.Content)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestEmailUsecase(t *testing.T) {
	suite.Run(t, new(EmailUsecaseTestSuite))
}
