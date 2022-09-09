package test

import (
	u "auth-plus-notification/cmd/usecases"
	t "auth-plus-notification/test/unit/mocks"

	"errors"
	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PushNotificationUsecaseTestSuite struct {
	suite.Suite
}

func (suite *PushNotificationUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	firebaseMocked := new(t.FirebaseMocked)
	firebaseMocked.On("SendPN", mockData.Email, mockData.Content).Return(true, nil)

	const number = 0.7
	randomManager := new(t.RandomPushNotificationManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(firebaseMocked, nil)

	pnUsecase := u.NewPushNotificationUsecase(randomManager)
	resp, err := pnUsecase.Send(mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func (suite *PushNotificationUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	firebaseMocked := new(t.FirebaseMocked)
	firebaseMocked.On("SendPN", mockData.Email, mockData.Content).Return(false, errors.New("failed"))

	const number = 0.7
	randomManager := new(t.RandomPushNotificationManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(firebaseMocked, nil)

	pnUsecase := u.NewPushNotificationUsecase(randomManager)
	resp, err := pnUsecase.Send(mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), resp, false)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestPushNotificationUsecase(t *testing.T) {
	suite.Run(t, new(PushNotificationUsecaseTestSuite))
}
