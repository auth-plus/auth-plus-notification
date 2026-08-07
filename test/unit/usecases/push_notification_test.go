package test

import (
	"github.com/stretchr/testify/mock"
	"context"
	u "auth-plus-notification/internal/usecases"
	d "auth-plus-notification/internal/usecases/driven"
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
	firebaseMocked.On("SendPN", mock.Anything, mockData.DeviceID, mockData.Title, mockData.Content).Return(nil)

	const number = 0.7
	randomManager := new(t.ManagerMocked[d.SendingPushNotification, float64])
	randomManager.On("GetInput", mock.Anything).Return(number, nil)
	randomManager.On("ChooseProvider", mock.Anything, number).Return(firebaseMocked, nil)

	pnUsecase := u.NewPushNotificationUsecase(randomManager)
	err := pnUsecase.Send(context.Background(), mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *PushNotificationUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	firebaseMocked := new(t.FirebaseMocked)
	firebaseMocked.On("SendPN", mock.Anything, mockData.DeviceID, mockData.Title, mockData.Content).Return(errors.New("failed"))

	const number = 0.7
	randomManager := new(t.ManagerMocked[d.SendingPushNotification, float64])
	randomManager.On("GetInput", mock.Anything).Return(number, nil)
	randomManager.On("ChooseProvider", mock.Anything, number).Return(firebaseMocked, nil)

	pnUsecase := u.NewPushNotificationUsecase(randomManager)
	err := pnUsecase.Send(context.Background(), mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestPushNotificationUsecase(t *testing.T) {
	suite.Run(t, new(PushNotificationUsecaseTestSuite))
}
