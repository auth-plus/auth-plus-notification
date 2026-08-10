package usecases

import (
	"auth-plus-notification/config"
	d "auth-plus-notification/internal/usecases/driven"
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// TelegramUsecase dependencies
type TelegramUsecase struct {
	manager d.Manager[d.SendingTelegram, float64]
	logger  *otelzap.Logger
}

// NewTelegramUsecase for instanciate a Telegram usecase
func NewTelegramUsecase(manager d.Manager[d.SendingTelegram, float64]) *TelegramUsecase {
	instance := new(TelegramUsecase)
	instance.manager = manager
	instance.logger = config.GetLogger()
	return instance
}

// Send method for sending an telegram message by using manager on dependecy
func (e *TelegramUsecase) Send(ctx context.Context, chatID int64, text string) error {
	e.logger.Ctx(ctx).Info("Sending Telegram message", zap.Int64("chatID", chatID))

	number, errI := e.manager.GetInput(ctx)
	if errI != nil {
		e.logger.Ctx(ctx).Error("Failed to get input for Telegram manager", zap.Error(errI))
		return errI
	}
	provider, errC := e.manager.ChooseProvider(ctx, number)
	if errC != nil {
		e.logger.Ctx(ctx).Error("Failed to choose Telegram provider", zap.Error(errC))
		return errC
	}
	err := provider.SendTele(ctx, chatID, text)
	if err != nil {
		e.logger.Ctx(ctx).Error("Failed to send Telegram message", zap.Error(err), zap.Int64("chatID", chatID))
		return err
	}

	e.logger.Ctx(ctx).Info("Telegram message sent successfully", zap.Int64("chatID", chatID))
	return nil
}
