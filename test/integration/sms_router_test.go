package test

import (
	app "auth-plus-notification/api/http"
	routes "auth-plus-notification/api/http/routes"
	"auth-plus-notification/config"
	mock "auth-plus-notification/test/mocks"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSmsHandler(t *testing.T) {
	r := app.Server()
	env := config.GetEnv()
	mockData := mock.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}
	payload := routes.SmsRequestBody{
		Phone:   mockData.Phone,
		Content: mockData.Content,
	}
	jsonValue, _ := json.Marshal(payload)

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://onesignal.com/api/v1").
		MatchHeader("Authorization", fmt.Sprintf("Basic %s", env.Providers.Onesignal.APIKey)).
		Post("/notifications").
		Reply(200)

	req, _ := http.NewRequest("POST", "/sms", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, "Ok", string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
