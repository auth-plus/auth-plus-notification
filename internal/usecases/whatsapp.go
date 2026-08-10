package usecases

import (
	"auth-plus-notification/config"
	d "auth-plus-notification/internal/usecases/driven"
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// WhatsappUsecase dependencies
type WhatsappUsecase struct {
	manager d.Manager[d.SendingWhatsapp, float64]
	logger  *otelzap.Logger
}

// NewWhatsappUsecase for instanciate a whatsapp usecase
func NewWhatsappUsecase(manager d.Manager[d.SendingWhatsapp, float64]) *WhatsappUsecase {
	instance := new(WhatsappUsecase)
	instance.manager = manager
	instance.logger = config.GetLogger()
	return instance
}

// Send method for sending an whatsapp by using manager on dependecy
func (e *WhatsappUsecase) Send(ctx context.Context, phone string, content string) error {
	e.logger.Ctx(ctx).Info("Sending WhatsApp message", zap.String("phone", phone))

	number, errI := e.manager.GetInput(ctx)
	if errI != nil {
		e.logger.Ctx(ctx).Error("Failed to get input for WhatsApp manager", zap.Error(errI))
		return errI
	}
	provider, errC := e.manager.ChooseProvider(ctx, number)
	if errC != nil {
		e.logger.Ctx(ctx).Error("Failed to choose WhatsApp provider", zap.Error(errC))
		return errC
	}
	err := provider.SendWhats(ctx, phone, content)
	if err != nil {
		e.logger.Ctx(ctx).Error("Failed to send WhatsApp message", zap.Error(err), zap.String("phone", phone))
		return err
	}

	e.logger.Ctx(ctx).Info("WhatsApp message sent successfully", zap.String("phone", phone))
	return nil
}
