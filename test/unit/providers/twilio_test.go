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

type TwilioTestSuite struct {
	suite.Suite
}

// see https://www.twilio.com/docs/openapi/using-twilio-postman-collections
func (suite *TwilioTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	env := config.GetEnv()
	defer gock.Off()
	gock.Observe(gock.DumpRequest)
	gock.New("https://api.twilio.com").
		Post(fmt.Sprintf("/2010-04-01/Accounts/%s/Messages.json", env.Providers.Twilio.AccountID)).
		Reply(200).
		JSON(map[string]interface{}{
			"account_sid":           "ACF8",
			"api_version":           "exercitation eiusmod amet laboris",
			"body":                  "consectetur exercitation officia magna",
			"date_created":          "consectetur exercitation proident",
			"date_sent":             "culpa sint labore",
			"date_updated":          "proident",
			"direction":             "outbound-call",
			"error_code":            -64742646,
			"error_message":         "nostrud sint Ut",
			"from":                  "Excepteur cillum aute",
			"messaging_service_sid": "MGDB",
			"num_media":             "dolor veniam non",
			"num_segments":          "et ut sint ut amet",
			"price":                 "Lorem reprehenderit",
			"price_unit":            "Excepteur tempor i",
			"sid":                   "MM79",
			"status":                "received",
			"to":                    "Ut non id",
			"uri":                   "magna sit esse",
		})

	provider := p.NewTwilio()
	err := provider.SendWhats(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func (suite *TwilioTestSuite) Test_fail_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	env := config.GetEnv()
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://api.twilio.com").
		Post(fmt.Sprintf("/2010-04-01/Accounts/%s/Messages.json", env.Providers.Twilio.AccountID)).
		Reply(401).
		JSON(map[string]interface{}{
			"code":      20003,
			"message":   "Authenticate",
			"more_info": "https://www.twilio.com/docs/errors/20003",
			"status":    401,
		})

	provider := p.NewTwilio()
	err := provider.SendWhats(mockData.Phone, mockData.Content)
	assert.Equal(suite.T(), err.Error(), "TwilioProvider: something went wrong")
}

func TestTwilio(t *testing.T) {
	suite.Run(t, new(TwilioTestSuite))
}
