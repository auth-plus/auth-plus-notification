// Package usecases contains all usecases
package usecases

import (
	"auth-plus-notification/config"
	m "auth-plus-notification/internal/managers"
	se "auth-plus-notification/internal/usecases/driven"
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// EmailUsecase dependencies
type EmailUsecase struct {
	manager se.Manager[se.SendingEmail, m.IPWarmingInput]
	logger  *otelzap.Logger
}

// NewEmailUsecase for instanciate a email usecase
func NewEmailUsecase(manager se.Manager[se.SendingEmail, m.IPWarmingInput]) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.manager = manager
	instance.logger = config.GetLogger()
	return instance
}

// Send method for sending an email by using manager on dependecy
func (e *EmailUsecase) Send(ctx context.Context, email string, subject string, content string) error {
	e.logger.Ctx(ctx).Info("Sending email", zap.String("email", email), zap.String("subject", subject))

	input, errI := e.manager.GetInput(ctx)
	if errI != nil {
		e.logger.Ctx(ctx).Error("Failed to get input for email manager", zap.Error(errI))
		return errI
	}
	provider, errC := e.manager.ChooseProvider(ctx, input)
	if errC != nil {
		e.logger.Ctx(ctx).Error("Failed to choose email provider", zap.Error(errC))
		return errC
	}
	err := provider.SendEmail(ctx, email, subject, content)
	if err != nil {
		e.logger.Ctx(ctx).Error("Failed to send email", zap.Error(err), zap.String("email", email))
		return err
	}

	e.logger.Ctx(ctx).Info("Email sent successfully", zap.String("email", email))
	return nil
}
