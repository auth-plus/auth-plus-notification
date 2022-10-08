package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// OneSignal struct must contains all private property to work
type OneSignal struct {
	url   string
	token string
	appID string
}

// NewOneSignal for instanciate a onesignal provider
func NewOneSignal() *OneSignal {
	env := config.GetEnv()
	instance := new(OneSignal)
	instance.url = "https://onesignal.com/api/v1/notifications"
	instance.token = env.Providers.Onesignal.APIKey
	instance.appID = ""
	return instance
}

type oneSignalEmailPayload struct {
	AppID   string    `json:"app_id"`
	Ids     [1]string `json:"include_player_ids"`
	Subject string    `json:"email_subject"`
	Body    string    `json:"email_body"`
}

// SendEmail implementation of SendingEmail: https://documentation.onesignal.com/reference/email-channel-properties
func (e *OneSignal) SendEmail(email string, subject string, content string) error {
	idList := [1]string{email}
	emailPayload := oneSignalEmailPayload{e.appID, idList, subject, content}

	json, jsonErr := json.Marshal(emailPayload)
	if jsonErr != nil {
		return jsonErr
	}
	reqErr := e.sendRequest(json)
	if reqErr != nil {
		return reqErr
	}
	return nil
}

type oneSignalPNPayloadContent struct {
	En string `json:"en"`
}

type oneSignalPNPayload struct {
	AppID   string                    `json:"app_id"`
	Ids     [1]string                 `json:"include_player_ids"`
	Data    map[string]interface{}    `json:"data"`
	Content oneSignalPNPayloadContent `json:"contents"`
}

// SendPN implementation of SendingPushNotification: https://documentation.onesignal.com/reference/push-channel-properties
func (e *OneSignal) SendPN(deviceID string, title string, content string) error {
	idList := [1]string{deviceID}
	pnPayload := oneSignalPNPayload{
		AppID:   e.appID,
		Ids:     idList,
		Data:    map[string]interface{}{"foo": "bar"},
		Content: oneSignalPNPayloadContent{En: content},
	}

	json, jsonErr := json.Marshal(pnPayload)
	if jsonErr != nil {
		return jsonErr
	}
	reqErr := e.sendRequest(json)
	if reqErr != nil {
		return reqErr
	}
	return nil
}

type oneSignalSMSPayload struct {
	AppID   string                    `json:"app_id"`
	Phone   [1]string                 `json:"include_phone_numbers"`
	Name    string                    `json:"name"`
	Content oneSignalPNPayloadContent `json:"contents"`
}

// SendSms implementation of SendingSms: https://documentation.onesignal.com/reference/sms-channel-properties
func (e *OneSignal) SendSms(phone string, content string) error {
	rand.Seed(time.Now().UnixNano())
	idList := [1]string{phone}
	name := fmt.Sprintf("%f", rand.Float64())
	smsPayload := oneSignalSMSPayload{
		AppID:   e.appID,
		Phone:   idList,
		Name:    name,
		Content: oneSignalPNPayloadContent{En: content},
	}

	json, jsonErr := json.Marshal(smsPayload)
	if jsonErr != nil {
		return jsonErr
	}
	reqErr := e.sendRequest(json)
	if reqErr != nil {
		return reqErr
	}
	return nil
}

func (e *OneSignal) sendRequest(json []byte) error {
	req, errReq := http.NewRequest("POST", e.url, bytes.NewBuffer(json))
	if errReq != nil {
		return errReq
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", e.token))
	req.Header.Add("content-type", "application/json; charset=utf-8")

	res, errHTTP := http.DefaultClient.Do(req)
	if errHTTP != nil {
		return errHTTP
	}

	defer res.Body.Close()
	_, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		return errBody
	}
	return nil
}
