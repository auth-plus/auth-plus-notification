package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type PushNotificationUsecase struct {
	sendingPushNotification d.SendingPushNotification
}

func NewPushNotificationUsecase(sendingPushNotification d.SendingPushNotification) *PushNotificationUsecase {
	instance := new(PushNotificationUsecase)
	instance.sendingPushNotification = sendingPushNotification
	return instance
}

func (e *PushNotificationUsecase) Send(deviceId string, title string, content string) {
	e.sendingPushNotification.SendPN(deviceId, title, content)
}
