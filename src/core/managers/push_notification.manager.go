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
	choosedProvider := choosePushNotificationProvider(deviceId, title, content)
	switch choosedProvider {
	case "Firebase":
		e.firebase.SendPN(deviceId, title, content)
		break
	case "OneSignal":
		e.onesignal.SendPN(deviceId, title, content)
		break
	}

}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func choosePushNotificationProvider(deviceId string, title string, content string) EnumPushNotificationProvider {
	if rand.Float64() < 0.5 {
		return "OneSignal"
	}
	return "Firebase"
}
