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

type WhatsappUsecaseTestSuite struct {
	suite.Suite
}

func (suite *WhatsappUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	twilioMocked := new(t.TwilioMocked)
	twilioMocked.On("SendWhats", mockData.Phone, mockData.Content).Return(true, nil)

	const number = 0.7
	randomManager := new(t.RandomWhatsappManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(twilioMocked, nil)

	whatsappUsecase := u.NewWhatsappUsecase(randomManager)
	resp, err := whatsappUsecase.Send(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func (suite *WhatsappUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	twilioMocked := new(t.TwilioMocked)
	twilioMocked.On("SendWhats", mockData.Phone, mockData.Content).Return(false, errors.New("failed"))

	const number = 0.7
	randomManager := new(t.RandomWhatsappManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(twilioMocked, nil)

	whatsappUsecase := u.NewWhatsappUsecase(randomManager)
	resp, err := whatsappUsecase.Send(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestWhatsappUsecase(t *testing.T) {
	suite.Run(t, new(WhatsappUsecaseTestSuite))
}
