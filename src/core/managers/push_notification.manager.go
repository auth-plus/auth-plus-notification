package managers

import (
	d "auth-plus-notification/core/usecases/driven"
	"math/rand"
)

//Class for PushNotificationManager
type PushNotificationManager struct {
	firebase  d.SendingPushNotification
	onesignal d.SendingPushNotification
}

func NewPushNotificationManager(firebase d.SendingPushNotification, onesignal d.SendingPushNotification) *PushNotificationManager {
	return &PushNotificationManager{firebase: firebase, onesignal: onesignal}
}

func (e *PushNotificationManager) SendPN(deviceId string, title string, content string) {
	provider := e.choosePushNotificationProvider(deviceId, title, content)
	provider.SendPN(deviceId, title, content)
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func (e *PushNotificationManager) choosePushNotificationProvider(deviceId string, title string, content string) d.SendingPushNotification {
	if rand.Float64() < 0.5 {
		return e.onesignal
	}
	return e.firebase
}
