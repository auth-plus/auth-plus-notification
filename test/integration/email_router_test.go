package test

import (
	app "auth-plus-notification/api/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	routes "auth-plus-notification/api/http/routes"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

type MockedData struct {
	Email   string `faker:"email"`
	Content string `faker:"sentence"`
}

func TestHomepageHandler(t *testing.T) {
	r := app.Server()
	mockData := MockedData{}
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}
	payload := routes.EmailRequestBody{
		Email:   mockData.Email,
		Content: mockData.Content,
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/email", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, "Ok", string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
