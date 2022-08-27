package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sendgrid struct {
	url   string
	token string
}

type SendgridEmailPayload struct {
	Personalizations string `json:"name"`
	From             string `json:"from"`
	Subject          string `json:"subject"`
	Content          string `json:"content"`
}

func NewSendgrid() *Sendgrid {
	instance := new(Sendgrid)
	env := config.GetEnv()
	instance.url = env.Providers.Sendgrid.Url
	instance.token = env.Providers.Sendgrid.ApiKey
	return instance
}

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
