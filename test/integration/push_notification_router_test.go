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

func TestPushNotificationHandler(t *testing.T) {
	r := app.Server()
	env := config.GetEnv()
	mockData := mock.MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}
	const deviceID string = "eqEiHBNFOisH1w2wlmZgUH:APA91bGV0clbAlzqgzgRJOFXawN1Kp4163nwjRSWZ46pNXnfBbwEiaV-Bop4R2YrnliXPMic3XXa9N7DU_orVMPjc7bPrgJcGk9bqwRzJYfR62J5fGdg_umaeSyHTyyTQeDGhRoBSsjR"
	payload := routes.PushNotificationRequestBody{
		DeviceID: deviceID,
		Title:    mockData.Title,
		Content:  mockData.Content,
	}
	jsonValue, _ := json.Marshal(payload)

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://onesignal.com/api/v1").
		MatchHeader("Authorization", fmt.Sprintf("Basic %s", env.Providers.Onesignal.APIKey)).
		Post("/notifications").
		Reply(200)

	req, _ := http.NewRequest("POST", "/push_notification", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, "Ok", string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
