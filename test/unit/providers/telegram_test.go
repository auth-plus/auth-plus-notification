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

type TelegramTestSuite struct {
	suite.Suite
}

func (suite *TelegramTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	env := config.GetEnv()
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://api.telegram.org").
		Post(fmt.Sprintf("/bot%s/getMe", env.Providers.Telegram.APIKey)).
		Reply(200).
		JSON(map[string]interface{}{
			"ok": true,
			"result": map[string]interface{}{
				"id":                          5198414170,
				"is_bot":                      true,
				"first_name":                  "Echo",
				"username":                    "EeCcHh0oBot",
				"can_join_groups":             true,
				"can_read_all_group_messages": false,
				"supports_inline_queries":     false,
			}})
	gock.New("https://api.telegram.org").
		Post(fmt.Sprintf("/bot%s/sendMessage", env.Providers.Telegram.APIKey)).
		Reply(200).
		JSON(map[string]interface{}{
			"ok": true,
			"result": map[string]interface{}{
				"message_id": 1498106028,
				"date":       1665100122575,
				"chat": map[string]interface{}{
					"id":   mockData.ChatID,
					"type": "private",
				},
			}})

	provider := p.NewTelegram()
	err := provider.SendTele(mockData.ChatID, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestTelegram(t *testing.T) {
	suite.Run(t, new(TelegramTestSuite))
}
