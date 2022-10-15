// Package providers contains all implementations of providers
package providers

import (
	"auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Firebase struct must contains all private property to work
type Firebase struct {
	client *http.Client
	logger *zap.Logger
}

// NewFirebase for instanciate a firebase provider
func NewFirebase() *Firebase {
	instance := new(Firebase)
	client, err := google.DefaultClient(oauth2.NoContext,
		"https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		log.Fatal(err)
	}
	instance.client = client
	instance.logger = config.GetLogger()
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

type errorType struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Status  string                   `json:"status"`
	Details []map[string]interface{} `json:"details"`
}
type errorRequest struct {
	Error errorType `json:"error"`
}

// SendPN implementation of SendingPushNotification
func (e *Firebase) SendPN(deviceID string, title string, content string) error {
	message := payload{
		Message: notification{
			Notification: notificationContent{
				Title: title,
				Body:  content,
			},
			Token: deviceID,
		},
	}
	jsonData, errJSON := json.Marshal(message)
	if errJSON != nil {
		return errJSON
	}

	resp, errReq := e.client.Post("https://fcm.googleapis.com/v1/projects/auth-plus-c2b74/messages:send", "application/json", bytes.NewBuffer(jsonData))
	if errReq != nil {
		return errReq
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := e.getError(resp)
		if err != nil {
			e.logger.Error(err.Error())
		}
		e.logger.Error(errMsg)
		return errors.New("FirebaseProvider: something went wrong")
	}
	return nil
}

func (e *Firebase) getError(resp *http.Response) (string, error) {
	respBody, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return "", errBody
	}
	var respJSON errorRequest
	errJSON := json.Unmarshal(respBody, &respJSON)
	if errJSON != nil {
		return "", errJSON
	}
	return respJSON.Error.Message, nil
}
