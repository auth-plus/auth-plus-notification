package managers

import (
	d "auth-plus-notification/internal/usecases/driven"
	"math/rand"
	"time"
)

// RandomTelegramManager must contains all provider that could be choosen
type RandomTelegramManager struct {
	telegram d.SendingTelegram
}

// NewRandomTelegramManager is a function for intanciate a pointer for PushNotification
func NewRandomTelegramManager(telegram d.SendingTelegram) *RandomTelegramManager {
	instance := new(RandomTelegramManager)
	instance.telegram = telegram
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *RandomTelegramManager) ChooseProvider(number float64) (d.SendingTelegram, error) {
	return e.telegram, nil
}

// GetInput is a function that generate a random number
func (e *RandomTelegramManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
