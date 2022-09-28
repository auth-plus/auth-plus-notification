package test

import (
	app "auth-plus-notification/api/http"
	routes "auth-plus-notification/api/http/routes"
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
)

func TestPushNotificationHandler(t *testing.T) {
	r := app.Server()
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
	req, _ := http.NewRequest("POST", "/push_notification", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, "Ok", string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
