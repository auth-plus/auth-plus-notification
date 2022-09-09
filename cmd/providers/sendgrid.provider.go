package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Sendgrid struct must contains all private property to work
type Sendgrid struct {
	url   string
	token string
}

// SendgridEmailPayload is the payload that sendgrid require
type SendgridEmailPayload struct {
	Personalizations string `json:"name"`
	From             string `json:"from"`
	Subject          string `json:"subject"`
	Content          string `json:"content"`
}

// NewSendgrid for instanciate a sendgrid provider
func NewSendgrid() *Sendgrid {
	instance := new(Sendgrid)
	env := config.GetEnv()
	instance.url = env.Providers.Sendgrid.URL
	instance.token = env.Providers.Sendgrid.APIKey
	return instance
}

// SendEmail implementation of SendingEmail
func (e *Sendgrid) SendEmail(email string, content string) (bool, error) {
	client := &http.Client{}
	emailPayload := SendgridEmailPayload{
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
