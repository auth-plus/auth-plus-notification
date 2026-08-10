package usecases

import (
	"auth-plus-notification/config"
	d "auth-plus-notification/internal/usecases/driven"
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// SmsUsecase dependencies
type SmsUsecase struct {
	manager d.Manager[d.SendingSms, float64]
	logger  *otelzap.Logger
}

// NewSmsUsecase for instanciate a sms usecase
func NewSmsUsecase(manager d.Manager[d.SendingSms, float64]) *SmsUsecase {
	instance := new(SmsUsecase)
	instance.manager = manager
	instance.logger = config.GetLogger()
	return instance
}

// Send method for sending an sms by using manager on dependecy
func (e *SmsUsecase) Send(ctx context.Context, phone string, content string) error {
	e.logger.Ctx(ctx).Info("Sending SMS", zap.String("phone", phone))

	number, errI := e.manager.GetInput(ctx)
	if errI != nil {
		e.logger.Ctx(ctx).Error("Failed to get input for SMS manager", zap.Error(errI))
		return errI
	}
	provider, errC := e.manager.ChooseProvider(ctx, number)
	if errC != nil {
		e.logger.Ctx(ctx).Error("Failed to choose SMS provider", zap.Error(errC))
		return errC
	}
	err := provider.SendSms(ctx, phone, content)
	if err != nil {
		e.logger.Ctx(ctx).Error("Failed to send SMS", zap.Error(err), zap.String("phone", phone))
		return err
	}

	e.logger.Ctx(ctx).Info("SMS sent successfully", zap.String("phone", phone))
	return nil
}
