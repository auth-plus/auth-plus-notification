package providers

import (
	config "auth-plus-notification/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Mailgun struct {
	url   string
	token string
}

type MailgunEmailPayload struct {
	Personalizations string `json:"name"`
	From             string `json:"from"`
	Subject          string `json:"subject"`
	Content          string `json:"content"`
}

func NewMailgun() *Mailgun {
	instance := new(Mailgun)
	env := config.GetEnv()
	instance.url = env.Providers.Mailgun.Url
	instance.token = env.Providers.Mailgun.ApiKey
	return instance
}

func (e *Mailgun) SendEmail(email string, content string) {
	client := &http.Client{}
	emailPayload := MailgunEmailPayload{
		Personalizations: "",
		From:             "",
		Subject:          "",
		Content:          "",
	}
	json, err := json.Marshal(emailPayload)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", e.url, bytes.NewBuffer(json))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "Bearer "+e.token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(f))
}
