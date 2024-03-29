package usecases

import (
	d "auth-plus-notification/internal/usecases/driven"
)

// PushNotificationUsecase dependencies
type PushNotificationUsecase struct {
	manager d.Manager[d.SendingPushNotification, float64]
}

// NewPushNotificationUsecase for instanciate a push notification usecase
func NewPushNotificationUsecase(manager d.Manager[d.SendingPushNotification, float64]) *PushNotificationUsecase {
	instance := new(PushNotificationUsecase)
	instance.manager = manager
	return instance
}

// Send method for sending an push notification by using manager on dependecy
func (e *PushNotificationUsecase) Send(deviceID string, title string, content string) error {
	number, errI := e.manager.GetInput()
	if errI != nil {
		return errI
	}
	provider, errC := e.manager.ChooseProvider(number)
	if errC != nil {
		return errC
	}
	return provider.SendPN(deviceID, title, content)
}
