package managers

import (
	d "auth-plus-notification/core/usecases/driven"
)

//Class for PushNotificationManager
type PushNotificationManager struct {
	firebase d.SendingPushNotification
}

func NewPushNotificationManager(firebase d.SendingPushNotification) *PushNotificationManager {
	return &PushNotificationManager{firebase: firebase}
}

func (e *PushNotificationManager) SendPN(deviceId string, title string, content string) {
	choosedProvider := choosePushNotificationProvider(deviceId, title, content)
	switch choosedProvider {
	case "Firebase":
		e.firebase.SendPN(deviceId, title, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func choosePushNotificationProvider(deviceId string, title string, content string) EnumPushNotificationProvider {
	return "Firebase"
}
