package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Sendgrid struct must contains all private property to work
type Sendgrid struct {
	url   string
	token string
}

// NewSendgrid for instanciate a sendgrid provider
func NewSendgrid() *Sendgrid {
	instance := new(Sendgrid)
	env := config.GetEnv()
	instance.url = "https://api.sendgrid.com/v3/mail/send"
	instance.token = env.Providers.Sendgrid.APIKey
	return instance
}

type sendgridEmailPayload struct {
	Personalizations [1]map[string]interface{} `json:"personalizations"`
	From             map[string]interface{}    `json:"from"`
	Subject          string                    `json:"subject"`
	Content          [1]map[string]interface{} `json:"content"`
}

// SendEmail implementation of SendingEmail
func (e *Sendgrid) SendEmail(email string, subject string, content string) error {
	client := &http.Client{}
	to := [1]map[string]interface{}{{
		"email": email,
	}}
	emailPayload := sendgridEmailPayload{
		Personalizations: [1]map[string]interface{}{{
			"to": to,
		}},
		From: map[string]interface{}{
			"email": "no-reply@auth-plus.app",
			"name":  "No Reply",
		},
		Subject: subject,
		Content: [1]map[string]interface{}{{
			"type":  "text/html",
			"value": content,
		}},
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
	resp, errExec := client.Do(req)
	if errExec != nil {
		return errExec
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := e.getError(resp)
		if err != nil {
			log.Println("Error parsing", err)
		}
		log.Println(errMsg)
		return errors.New("OneSignalProvider: something went wrong")
	}

	return nil
}

func (e *Sendgrid) getError(resp *http.Response) (string, error) {
	respBody, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return "", errBody
	}
	return string(respBody), nil
}
