package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
	"time"
)

type randomTelegramManager struct {
	telegram d.SendingTelegram
}

// NewRandomTelegramManager is a function for intanciate a pointer for PushNotification
func NewRandomTelegramManager(telegram d.SendingTelegram) *randomTelegramManager {
	instance := new(randomTelegramManager)
	instance.telegram = telegram
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *randomTelegramManager) ChooseProvider(number float64) (d.SendingTelegram, error) {
	return e.telegram, nil
}

// GetInput is a function that generate a random number
func (e *randomTelegramManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
