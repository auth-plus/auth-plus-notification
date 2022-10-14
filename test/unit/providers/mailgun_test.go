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

type MailgunTestSuite struct {
	suite.Suite
}

func (suite *MailgunTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	env := config.GetEnv()

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://api.mailgun.net").
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", env.Providers.Mailgun.APIKey)).
		Post("/").
		Reply(200)

	provider := p.NewMailgun()
	err := provider.SendEmail(mockData.Email, mockData.Subject, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *MailgunTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	env := config.GetEnv()

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://api.mailgun.net").
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", env.Providers.Mailgun.APIKey)).
		Post("/").
		Reply(500).
		BodyString(mockData.Error)

	provider := p.NewMailgun()
	err := provider.SendEmail(mockData.Email, mockData.Subject, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "MailgunProvider: something went wrong")
}

func TestMailgun(t *testing.T) {
	suite.Run(t, new(MailgunTestSuite))
}
