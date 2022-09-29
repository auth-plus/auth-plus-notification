package test

import (
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

type WhatsappUsecaseTestSuite struct {
	suite.Suite
}

func (suite *WhatsappUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	twilioMocked := new(t.TwilioMocked)
	twilioMocked.On("SendWhats", mockData.Phone, mockData.Content).Return(nil)

	const number = 0.7
	randomManager := new(t.ManagerMocked[d.SendingWhatsapp])
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(twilioMocked, nil)

	whatsappUsecase := u.NewWhatsappUsecase(randomManager)
	err := whatsappUsecase.Send(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *WhatsappUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	twilioMocked := new(t.TwilioMocked)
	twilioMocked.On("SendWhats", mockData.Phone, mockData.Content).Return(errors.New("failed"))

	const number = 0.7
	randomManager := new(t.ManagerMocked[d.SendingWhatsapp])
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(twilioMocked, nil)

	whatsappUsecase := u.NewWhatsappUsecase(randomManager)
	err := whatsappUsecase.Send(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestWhatsappUsecase(t *testing.T) {
	suite.Run(t, new(WhatsappUsecaseTestSuite))
}
