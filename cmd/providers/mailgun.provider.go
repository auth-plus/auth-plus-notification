package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Mailgun struct must contains all private property to work
type Mailgun struct {
	url   string
	token string
}

// MailgunEmailPayload is the payload htat mailgun require
type MailgunEmailPayload struct {
	Personalizations string `json:"name"`
	From             string `json:"from"`
	Subject          string `json:"subject"`
	Content          string `json:"content"`
}

// NewMailgun for instanciate a mailgun provider
func NewMailgun() *Mailgun {
	instance := new(Mailgun)
	env := config.GetEnv()
	instance.url = env.Providers.Mailgun.URL
	instance.token = env.Providers.Mailgun.APIKey
	return instance
}

// SendEmail implementation of SendingEmail
func (e *Mailgun) SendEmail(email string, content string) (bool, error) {
	client := &http.Client{}
	emailPayload := MailgunEmailPayload{
		Personalizations: "",
		From:             "",
		Subject:          "",
		Content:          "",
	}
	json, err := json.Marshal(emailPayload)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", e.url, bytes.NewBuffer(json))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "Bearer "+e.token)
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	fmt.Println(string(f))
	return true, nil
}
