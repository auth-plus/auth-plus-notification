package usecases

import (
	d "auth-plus-notification/cmd/usecases/driven"
)

type PushNotificationUsecase struct {
	manager d.PushNotificatioManager
}

func NewPushNotificationUsecase(manager d.PushNotificatioManager) *PushNotificationUsecase {
	instance := new(PushNotificationUsecase)
	instance.manager = manager
	return instance
}

func (e *PushNotificationUsecase) Send(deviceId string, title string, content string) (bool, error) {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return false, errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return false, errC
	}
	return provider.SendPN(deviceId, title, content)
}
