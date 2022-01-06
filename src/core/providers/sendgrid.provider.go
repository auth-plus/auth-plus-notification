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

type Sendgrid struct {
	url   string
	token string
}

type EmailPayload struct {
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

func (e *Sendgrid) SendEmail(email string, content string) {
	client := &http.Client{}
	emailPayload := EmailPayload{
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
