package test

import (
	"auth-plus-notification/config"
	p "auth-plus-notification/internal/providers"
	t "auth-plus-notification/test/mocks"

	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/gock.v1"
)

type OnesignalTestSuite struct {
	suite.Suite
}

func (suite *OnesignalTestSuite) Test_succeed_when_sending_email() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	env := config.GetEnv()

	defer gock.Off() // Flush pending mocks after test execution
	gock.Observe(gock.DumpRequest)
	gock.New("https://onesignal.com/api/v1").
		MatchHeader("Authorization", fmt.Sprintf("Basic %s", env.Providers.Onesignal.APIKey)).
		Post("/notifications").
		Reply(200)

	provider := p.NewOneSignal()
	err := provider.SendEmail(mockData.Email, mockData.Subject, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}
func (suite *OnesignalTestSuite) Test_succeed_when_sending_push_notification() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	env := config.GetEnv()

	defer gock.Off() // Flush pending mocks after test execution
	gock.Observe(gock.DumpRequest)
	gock.New("https://onesignal.com/api/v1").
		MatchHeader("Authorization", fmt.Sprintf("Basic %s", env.Providers.Onesignal.APIKey)).
		Post("/notifications").
		Reply(200)

	provider := p.NewOneSignal()
	err := provider.SendPN(mockData.DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}
func (suite *OnesignalTestSuite) Test_succeed_when_sending_sms() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	env := config.GetEnv()

	defer gock.Off() // Flush pending mocks after test execution
	gock.Observe(gock.DumpRequest)
	gock.New("https://onesignal.com/api/v1").
		MatchHeader("Authorization", fmt.Sprintf("Basic %s", env.Providers.Onesignal.APIKey)).
		Post("/notifications").
		Reply(200)

	provider := p.NewOneSignal()
	err := provider.SendSms(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestOnesignal(t *testing.T) {
	suite.Run(t, new(OnesignalTestSuite))
}
