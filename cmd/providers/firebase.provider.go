// Package providers contains all implementations of providers
package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Firebase struct must contains all private property to work
type Firebase struct {
	client *http.Client
}

// NewFirebase for instanciate a firebase provider
func NewFirebase() *Firebase {
	instance := new(Firebase)
	// env := config.GetEnv()
	client, err := google.DefaultClient(oauth2.NoContext,
		"https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		log.Fatal(err)
	}
	instance.client = client
	return instance
}

type notificationContent struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type notification struct {
	Notification notificationContent `json:"notification"`
	Token        string              `json:"token"`
}

type payload struct {
	Message notification `json:"message"`
}

// SendPN implementation of SendingPushNotification
func (e *Firebase) SendPN(deviceID string, title string, content string) error {
	// See documentation on defining a message payload.
	message := payload{
		Message: notification{
			Notification: notificationContent{
				Title: title,
				Body:  content,
			},
			Token: deviceID,
		},
	}
	jsonData, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}
	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := e.client.Post("https://fcm.googleapis.com/v1/projects/auth-plus-c2b74/messages:send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	return nil
}
