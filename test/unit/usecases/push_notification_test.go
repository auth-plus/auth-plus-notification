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

type PushNotificationUsecaseTestSuite struct {
	suite.Suite
}

func (suite *PushNotificationUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	firebaseMocked := new(t.FirebaseMocked)
	firebaseMocked.On("SendPN", mockData.DeviceID, mockData.Title, mockData.Content).Return(nil)

	const number = 0.7
	randomManager := new(t.RandomPushNotificationManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(firebaseMocked, nil)

	pnUsecase := u.NewPushNotificationUsecase(randomManager)
	err := pnUsecase.Send(mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *PushNotificationUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	firebaseMocked := new(t.FirebaseMocked)
	firebaseMocked.On("SendPN", mockData.DeviceID, mockData.Title, mockData.Content).Return(errors.New("failed"))

	const number = 0.7
	randomManager := new(t.RandomPushNotificationManagerMocked)
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(firebaseMocked, nil)

	pnUsecase := u.NewPushNotificationUsecase(randomManager)
	err := pnUsecase.Send(mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestPushNotificationUsecase(t *testing.T) {
	suite.Run(t, new(PushNotificationUsecaseTestSuite))
}
