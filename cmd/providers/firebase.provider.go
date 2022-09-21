// Package providers contains all implementations of providers
package providers

import (
	"auth-plus-notification/config"
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// Firebase struct must contains all private property to work
type Firebase struct {
	app *firebase.App
}

// NewFirebase for instanciate a firebase provider
func NewFirebase() *Firebase {
	instance := new(Firebase)
	env := config.GetEnv()
	opt := option.WithCredentialsFile(env.Providers.Firebase.Credential)
	config := &firebase.Config{ProjectID: env.Providers.Firebase.AppName}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	instance.app = app
	return instance
}

// SendPN implementation of SendingPushNotification
func (e *Firebase) SendPN(deviceID string, title string, content string) (bool, error) {
	ctx := context.Background()
	client, err := e.app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  content,
		},
		Token: deviceID,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	return true, nil
}
