package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
	"time"
)

type randomSmsManager struct {
	sns       d.SendingSms
	onesignal d.SendingSms
}

// NewRandomSmsManager is a function for intanciate a pointer for PushNotification
func NewRandomSmsManager(sns d.SendingSms, onesignal d.SendingSms) *randomSmsManager {
	instance := new(randomSmsManager)
	instance.sns = sns
	instance.onesignal = onesignal
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *randomSmsManager) ChooseProvider(number float64) (d.SendingSms, error) {
	if number < 0.5 {
		return e.onesignal, nil
	}
	return e.sns, nil
}

// GetInput is a function that generate a random number
func (e *randomSmsManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
