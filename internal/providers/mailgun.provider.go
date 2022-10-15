package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

// Mailgun struct must contains all private property to work
type Mailgun struct {
	url    string
	token  string
	logger *zap.Logger
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

type mailgunEmailPayload struct {
	Personalizations string `json:"name"`
	From             string `json:"from"`
	Subject          string `json:"subject"`
	Content          string `json:"content"`
}

// SendEmail implementation of SendingEmail: https://documentation.mailgun.com/en/latest/api-intro.html#introduction
func (e *Mailgun) SendEmail(email string, subject string, content string) error {
	client := &http.Client{}
	emailPayload := mailgunEmailPayload{
		Personalizations: email,
		From:             "noreply@auth-plus.com",
		Subject:          subject,
		Content:          content,
	}

	json, errJSON := json.Marshal(emailPayload)
	if errJSON != nil {
		return errJSON
	}

	req, errReq := http.NewRequest("POST", e.url, bytes.NewBuffer(json))
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
			e.logger.Error(err.Error())
		}
		e.logger.Error(errMsg)
		return errors.New("MailgunProvider: something went wrong")
	}

	return nil
}

func (e *Mailgun) getError(resp *http.Response) (string, error) {
	respBody, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return "", errBody
	}
	return string(respBody), nil
}
