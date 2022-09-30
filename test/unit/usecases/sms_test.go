package test

import (
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

type SmsUsecaseTestSuite struct {
	suite.Suite
}

func (suite *SmsUsecaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	snsMocked := new(t.SnsMocked)
	snsMocked.On("SendSms", mockData.Phone, mockData.Content).Return(nil)

	const number = 0.7
	randomManager := new(t.ManagerMocked[d.SendingSms, float64])
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(snsMocked, nil)

	smsUsecase := u.NewSmsUsecase(randomManager)
	err := smsUsecase.Send(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *SmsUsecaseTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	snsMocked := new(t.SnsMocked)
	snsMocked.On("SendSms", mockData.Phone, mockData.Content).Return(errors.New("failed"))

	const number = 0.7
	randomManager := new(t.ManagerMocked[d.SendingSms, float64])
	randomManager.On("GetInput").Return(number, nil)
	randomManager.On("ChooseProvider", number).Return(snsMocked, nil)

	smsUsecase := u.NewSmsUsecase(randomManager)
	err := smsUsecase.Send(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "failed")
}

func TestSmsUsecase(t *testing.T) {
	suite.Run(t, new(SmsUsecaseTestSuite))
}
