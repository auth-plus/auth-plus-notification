package test

import (
	m "auth-plus-notification/cmd/managers"
	u "auth-plus-notification/cmd/usecases"
	d "auth-plus-notification/cmd/usecases/driven"
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
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	sendgridMocked := new(t.SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Content).Return(nil)

	input := m.IPWarmingInput{Sendgrid: 100, Mailgun: 50, Onesignal: 75}
	randomEmailManager := new(t.ManagerMocked[d.SendingEmail, m.IPWarmingInput])
	randomEmailManager.On("GetInput").Return(input, nil)
	randomEmailManager.On("ChooseProvider", input).Return(sendgridMocked, nil)

	emailUsecase := u.NewEmailUsecase(randomEmailManager)
	err := emailUsecase.Send(mockData.Email, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *EmailUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	sendgridMocked := new(t.SendgridMocked)
	sendgridMocked.On("SendEmail", mockData.Email, mockData.Content).Return(errors.New("failed"))

	input := m.IPWarmingInput{Sendgrid: 100, Mailgun: 50, Onesignal: 75}
	randomEmailManager := new(t.ManagerMocked[d.SendingEmail, m.IPWarmingInput])
	randomEmailManager.On("GetInput").Return(input, nil)
	randomEmailManager.On("ChooseProvider", input).Return(sendgridMocked, nil)

	emailUsecase := u.NewEmailUsecase(randomEmailManager)
	err := emailUsecase.Send(mockData.Email, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestEmailUsecase(t *testing.T) {
	suite.Run(t, new(EmailUsecaseTestSuite))
}
