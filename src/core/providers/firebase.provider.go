package providers

import (
	"context"
	"fmt"
	"log"

	config "auth-plus-notification/config"

	firebase "firebase.google.com/go"
	messaging "firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type Firebase struct {
	app      *firebase.App
	logopath string
}

func NewFirebase() *Firebase {
	instance := new(Firebase)
	env := config.GetEnv()
	opt := option.WithCredentialsFile(env.Providers.Firebase.CredentialPath)
	config := &firebase.Config{ProjectID: env.Providers.Firebase.AppName}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	instance.app = app
	instance.logopath = "https://upload.wikimedia.org/wikipedia/commons/c/ce/Twitter_Logo.png"
	return instance
}

func (e *Firebase) SendPN(deviceId string, title string, content string) {
	ctx := context.Background()
	client, err := e.app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     content,
			ImageURL: e.logopath,
		},
		Token: deviceId,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}
