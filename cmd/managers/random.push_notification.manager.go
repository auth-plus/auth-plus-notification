package managers

import (
	d "auth-plus-notification/cmd/usecases/driven"
	"math/rand"
	"time"
)

type randomPushNotificationManager struct {
	firebase  d.SendingPushNotification
	onesignal d.SendingPushNotification
}

// NewRandomPushNotificationManager is a function for intanciate a pointer for PushNotification
func NewRandomPushNotificationManager(firebase d.SendingPushNotification, onesignal d.SendingPushNotification) *randomPushNotificationManager {
	instance := new(randomPushNotificationManager)
	instance.firebase = firebase
	instance.onesignal = onesignal
	return instance
}

// ChooseProvider is a function for choosing a provider based on a number
func (e *randomPushNotificationManager) ChooseProvider(number float64) (d.SendingPushNotification, error) {
	if number < 0.5 {
		return e.onesignal, nil
	}
	return e.firebase, nil
}

// GetInput is a function that generate a random number
func (e *randomPushNotificationManager) GetInput() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
