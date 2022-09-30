package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	m "auth-plus-notification/cmd/managers"
	t "auth-plus-notification/test/mocks"
)

type RandomTelegramManagerTestSuite struct {
	suite.Suite
}

func (suite *RandomTelegramManagerTestSuite) Test_succeed_when_choosing_telegram() {
	telegramMocked := new(t.TelegramMocked)
	const number = 0.7
	smsManager := m.NewRandomTelegramManager(telegramMocked)
	provider, err := smsManager.ChooseProvider(number)
	assert.Equal(suite.T(), provider, telegramMocked)
	assert.Equal(suite.T(), err, nil)
}

func TestRandomTelegramManager(t *testing.T) {
	suite.Run(t, new(RandomTelegramManagerTestSuite))
}
