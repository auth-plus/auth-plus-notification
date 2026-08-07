package providers

import (
	"auth-plus-notification/config"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Sendgrid struct must contains all private property to work
type Sendgrid struct {
	url    string
	token  string
	logger *otelzap.Logger
}

// NewSendgrid for instanciate a sendgrid provider
func NewSendgrid() *Sendgrid {
	instance := new(Sendgrid)
	env := config.GetEnv()
	instance.url = "https://api.sendgrid.com/v3/mail/send"
	instance.token = env.Providers.Sendgrid.APIKey
	instance.logger = config.GetLogger()
	return instance
}

type sendgridBody struct {
	Personalizations [1]map[string]interface{} `json:"personalizations"`
	From             map[string]interface{}    `json:"from"`
	Content          [1]map[string]interface{} `json:"content"`
}

// SendEmail implementation of SendingEmail
func (e *Sendgrid) SendEmail(ctx context.Context, email string, subject string, content string) error {
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	to := [1]map[string]interface{}{{
		"email": email,
	}}

	personalizations := [1]map[string]interface{}{{
		"to":      to,
		"subject": subject,
	}}
	from := map[string]interface{}{
		"email": "no-reply@auth-plus.app",
		"name":  "No Reply",
	}

	contentObj := [1]map[string]interface{}{{
		"type":  "text/html",
		"value": content,
	}}

	body := sendgridBody{
		Personalizations: personalizations,
		From:             from,
		Content:          contentObj,
	}

	json, errJ := json.Marshal(body)
	if errJ != nil {
		return errJ
	}

	req, errReq := http.NewRequest("POST", e.url, bytes.NewBuffer(json))
	if errReq != nil {
		return errReq
	}
	req = req.WithContext(ctx)
	if errReq != nil {
		return errReq
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "Bearer "+e.token)
	resp, errExec := client.Do(req)
	if errExec != nil {
		return errExec
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := e.getError(resp)
		if err != nil {
			e.logger.Ctx(ctx).Error(err.Error())
		}
		e.logger.Ctx(ctx).Error(errMsg)
		return errors.New("SendgridProvider: something went wrong")
	}

	return nil
}

func (e *Sendgrid) getError(resp *http.Response) (string, error) {
	respBody, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return "", errBody
	}
	return string(respBody), nil
}
