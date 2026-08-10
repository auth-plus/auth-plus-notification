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

// Mailgun struct must contains all private property to work
type Mailgun struct {
	url    string
	token  string
	logger *otelzap.Logger
}

// NewMailgun for instanciate a mailgun provider
func NewMailgun() *Mailgun {
	instance := new(Mailgun)
	env := config.GetEnv()
	instance.url = "https://api.mailgun.net"
	instance.token = env.Providers.Mailgun.APIKey
	instance.logger = config.GetLogger()
	return instance
}

type mailgunBody struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	HTML    string `json:"html"`
}

// SendEmail implementation of SendingEmail: https://documentation.mailgun.com/en/latest/api-intro.html#introduction
func (e *Mailgun) SendEmail(ctx context.Context, email string, subject string, content string) error {

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	body := mailgunBody{
		To:      email,
		Subject: subject,
		From:    "noreply@auth-plus.com",
		HTML:    content,
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
		return errors.New("MailgunProvider: something went wrong")
	}

	return nil
}

func (e *Mailgun) getError(resp *http.Response) (string, error) {
	respBody, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return "", errBody
	}
	return string(respBody), nil
}
