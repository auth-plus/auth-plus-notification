package usecases

import (
	d "auth-plus-notification/core/usecases/driven"
)

type PushNotificationUsecase struct {
	sendingPushNotification d.SendingPushNotification
}

func NewPushNotificationUsecase() *PushNotificationUsecase {
	instance := new(PushNotificationUsecase)
	return instance
}

func (e *PushNotificationUsecase) Send(deviceId string, title string, content string) {
	e.sendingPushNotification.Send(deviceId, title, content)
}
