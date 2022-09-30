package test

import (
	p "auth-plus-notification/cmd/providers"
	t "auth-plus-notification/test/mocks"

	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/gock.v1"
)

type SNSTestSuite struct {
	suite.Suite
}

func (suite *SNSTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://sns.us-west-2.amazonaws.com").
		Post("/").
		Reply(200)

	provider := p.NewSNS()
	err := provider.SendSms(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestSNS(t *testing.T) {
	suite.Run(t, new(SNSTestSuite))
}
