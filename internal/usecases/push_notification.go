package usecases

import (
	"auth-plus-notification/config"
	d "auth-plus-notification/internal/usecases/driven"
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// PushNotificationUsecase dependencies
type PushNotificationUsecase struct {
	manager d.Manager[d.SendingPushNotification, float64]
	logger  *otelzap.Logger
}

// NewPushNotificationUsecase for instanciate a push notification usecase
func NewPushNotificationUsecase(manager d.Manager[d.SendingPushNotification, float64]) *PushNotificationUsecase {
	instance := new(PushNotificationUsecase)
	instance.manager = manager
	instance.logger = config.GetLogger()
	return instance
}

// Send method for sending an push notification by using manager on dependecy
func (e *PushNotificationUsecase) Send(ctx context.Context, deviceID string, title string, content string) error {
	e.logger.Ctx(ctx).Info("Sending push notification", zap.String("deviceID", deviceID), zap.String("title", title))

	number, errI := e.manager.GetInput(ctx)
	if errI != nil {
		e.logger.Ctx(ctx).Error("Failed to get input for push notification manager", zap.Error(errI))
		return errI
	}
	provider, errC := e.manager.ChooseProvider(ctx, number)
	if errC != nil {
		e.logger.Ctx(ctx).Error("Failed to choose push notification provider", zap.Error(errC))
		return errC
	}
	err := provider.SendPN(ctx, deviceID, title, content)
	if err != nil {
		e.logger.Ctx(ctx).Error("Failed to send push notification", zap.Error(err), zap.String("deviceID", deviceID))
		return err
	}

	e.logger.Ctx(ctx).Info("Push notification sent successfully", zap.String("deviceID", deviceID))
	return nil
}
