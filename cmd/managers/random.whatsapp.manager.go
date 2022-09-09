package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
	"time"
)

type randomWhatsappManager struct {
	twilio d.SendingWhatsapp
}

// NewRandomWhatsappManager is a function for intanciate a pointer for PushNotification
func NewRandomWhatsappManager(twilio d.SendingWhatsapp) *randomWhatsappManager {
	instance := new(randomWhatsappManager)
	instance.twilio = twilio
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *randomWhatsappManager) ChooseProvider(number float64) (d.SendingWhatsapp, error) {
	return e.twilio, nil
}

// GetInput is a function that generate a random number
func (e *randomWhatsappManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
