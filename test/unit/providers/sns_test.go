package test

import (
	p "auth-plus-notification/internal/providers"
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

type xmlError struct {
	Type    string `xml:"Type"`
	Code    string `xml:"Code"`
	Message string `xml:"Message"`
}

type xmlErrorResponse struct {
	Error xmlError `xml:"Error"`
}

type xml struct {
	ErrorResponse xmlErrorResponse `xml:"ErrorResponse"`
}

func (suite *SNSTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	defer gock.Off() // Flush pending mocks after test execution
	gock.Observe(gock.DumpRequest)
	gock.New("https://sns.us-west-2.amazonaws.com").
		Post("/").
		Reply(400).
		XML(xml{ErrorResponse: xmlErrorResponse{
			Error: xmlError{
				Type:    "Sender",
				Code:    "IncompleteSignature",
				Message: "Authorization header requires e...",
			},
		}})

	provider := p.NewSNS()
	err := provider.SendSms(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestSNS(t *testing.T) {
	suite.Run(t, new(SNSTestSuite))
}
