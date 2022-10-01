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

type SendgridTestSuite struct {
	suite.Suite
}

func (suite *SendgridTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	env := config.GetEnv()

	defer gock.Off() // Flush pending mocks after test execution
	gock.Observe(gock.DumpRequest)
	gock.New("https://api.sendgrid.com/v3/mail").
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", env.Providers.Sendgrid.APIKey)).
		Post("/send").
		Reply(200)

	provider := p.NewSendgrid()
	err := provider.SendEmail(mockData.Email, mockData.Subject, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestSendgrid(t *testing.T) {
	suite.Run(t, new(SendgridTestSuite))
}
