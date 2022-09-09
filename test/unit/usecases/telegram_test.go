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

type TelegramUsecaseTestSuite struct {
	suite.Suite
}

func (suite *TelegramUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	telegramMocked := new(t.TelegramMocked)
	telegramMocked.On("SendTele", mockData.ChatID, mockData.Content).Return(true, nil)

	const number = 0.7
	randomManager := new(t.RandomTelegramManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(telegramMocked, nil)

	telegramUsecase := u.NewTelegramUsecase(randomManager)
	resp, err := telegramUsecase.Send(mockData.ChatID, mockData.Content)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func (suite *TelegramUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	telegramMocked := new(t.TelegramMocked)
	telegramMocked.On("SendTele", mockData.ChatID, mockData.Content).Return(false, errors.New("failed"))

	const number = 0.7
	randomManager := new(t.RandomTelegramManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(telegramMocked, nil)

	telegramUsecase := u.NewTelegramUsecase(randomManager)
	resp, err := telegramUsecase.Send(mockData.ChatID, mockData.Content)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestTelegramUsecase(t *testing.T) {
	suite.Run(t, new(TelegramUsecaseTestSuite))
}
