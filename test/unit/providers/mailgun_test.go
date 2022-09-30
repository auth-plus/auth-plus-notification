package test

import (
	p "auth-plus-notification/cmd/providers"
	"auth-plus-notification/config"
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
	gock.Observe(gock.DumpRequest)
	gock.New(env.Providers.Mailgun.URL).
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", env.Providers.Mailgun.APIKey)).
		Post("/").
		Reply(200)

	provider := p.NewMailgun()
	err := provider.SendEmail(mockData.Email, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestMailgun(t *testing.T) {
	suite.Run(t, new(MailgunTestSuite))
}
