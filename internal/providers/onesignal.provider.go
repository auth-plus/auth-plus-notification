package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// OneSignal struct must contains all private property to work
type OneSignal struct {
	url    string
	token  string
	appID  string
	logger *otelzap.Logger
}

// NewOneSignal for instanciate a onesignal provider
func NewOneSignal() *OneSignal {
	env := config.GetEnv()
	instance := new(OneSignal)
	instance.url = "https://onesignal.com/api/v1/notifications"
	instance.token = env.Providers.Onesignal.APIKey
	instance.appID = ""
	instance.logger = config.GetLogger()
	return instance
}

type oneSignalEmailPayload struct {
	AppID   string    `json:"app_id"`
	IDs     [1]string `json:"include_player_ids"`
	Subject string    `json:"email_subject"`
	Body    string    `json:"email_body"`
}

// SendEmail implementation of SendingEmail: https://documentation.onesignal.com/reference/email-channel-properties
func (e *OneSignal) SendEmail(ctx context.Context, email string, subject string, content string) error {
	idList := [1]string{email}
	emailPayload := oneSignalEmailPayload{e.appID, idList, subject, content}

	json, jsonErr := json.Marshal(emailPayload)
	if jsonErr != nil {
		return jsonErr
	}
	reqErr := e.sendRequest(ctx, json)
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
	IDs     [1]string                 `json:"include_player_ids"`
	Data    map[string]interface{}    `json:"data"`
	Content oneSignalPNPayloadContent `json:"contents"`
}

// SendPN implementation of SendingPushNotification: https://documentation.onesignal.com/reference/push-channel-properties
func (e *OneSignal) SendPN(ctx context.Context, deviceID string, title string, content string) error {
	idList := [1]string{deviceID}
	pnPayload := oneSignalPNPayload{
		AppID:   e.appID,
		IDs:     idList,
		Data:    map[string]interface{}{"title": title, "content": content},
		Content: oneSignalPNPayloadContent{En: content},
	}

	json, jsonErr := json.Marshal(pnPayload)
	if jsonErr != nil {
		return jsonErr
	}
	reqErr := e.sendRequest(ctx, json)
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
func (e *OneSignal) SendSms(ctx context.Context, phone string, content string) error {
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
	reqErr := e.sendRequest(ctx, json)
	if reqErr != nil {
		return reqErr
	}
	return nil
}

func (e *OneSignal) sendRequest(ctx context.Context, json []byte) error {
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	req, errReq := http.NewRequest("POST", e.url, bytes.NewBuffer(json))
	if errReq != nil {
		return errReq
	}
	req = req.WithContext(ctx)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", e.token))
	req.Header.Add("content-type", "application/json; charset=utf-8")

	resp, errHTTP := client.Do(req)
	if errHTTP != nil {
		return errHTTP
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMsg, err := e.getError(resp)
		if err != nil {
			e.logger.Ctx(ctx).Error(err.Error())
		}
		e.logger.Ctx(ctx).Error(errMsg)
		return errors.New("OneSignalProvider: something went wrong")
	}
	return nil
}

func (e *OneSignal) getError(resp *http.Response) (string, error) {
	respBody, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return "", errBody
	}
	return string(respBody), nil
}
