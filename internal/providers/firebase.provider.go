// Package providers contains all implementations of providers
package providers

import (
	"auth-plus-notification/config"
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"google.golang.org/api/option"
)

// Firebase struct must contains all private property to work
type Firebase struct {
	app    *firebase.App
	logger *otelzap.Logger
}

// NewFirebase for instanciate a firebase provider
func NewFirebase() (*Firebase, error) {
	instance := new(Firebase)
	opt := option.WithCredentialsFile("./service-account-file.json")
	client, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	instance.app = client
	instance.logger = config.GetLogger()
	return instance, nil
}

// SendPN implementation of SendingPushNotification
func (e *Firebase) SendPN(ctx context.Context, deviceID string, title string, content string) error {
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"Title": title,
			"Body":  content,
		},
		Token: deviceID,
	}
	client, err := e.app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("error getting Messaging client: %v", err)
	}
	_, errReq := client.Send(ctx, message)
	if errReq != nil {
		return errReq
	}
	return nil
}
