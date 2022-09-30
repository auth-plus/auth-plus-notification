package managers

import (
	d "auth-plus-notification/internal/usecases/driven"
	"math/rand"
	"time"
)

// RandomWhatsappManager must contains all provider that could be choosen
type RandomWhatsappManager struct {
	twilio d.SendingWhatsapp
}

// NewRandomWhatsappManager is a function for intanciate a pointer for PushNotification
func NewRandomWhatsappManager(twilio d.SendingWhatsapp) *RandomWhatsappManager {
	instance := new(RandomWhatsappManager)
	instance.twilio = twilio
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *RandomWhatsappManager) ChooseProvider(number float64) (d.SendingWhatsapp, error) {
	return e.twilio, nil
}

// GetInput is a function that generate a random number
func (e *RandomWhatsappManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
