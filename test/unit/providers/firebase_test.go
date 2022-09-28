package test

import (
	p "auth-plus-notification/cmd/providers"
	t "auth-plus-notification/test/mocks"

	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/gock.v1"
)

type FirebaseTestSuite struct {
	suite.Suite
}

func (suite *FirebaseTestSuite) Test_succeed_when_sending() {
	mockData := t.MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}

	const DeviceID string = "cDxDrGiXRnMXFXFsyDLSY5:APA91bF9QI6-YU2eJr7JcF8u6lrAGAIBEpG4j3IOrU2h2EFUahYM1z0fMR_IyqybrOfc62ASy5uRyg1uzjR2trlpmZujQ79-QQAf7iSYj_4HZn3fWYg9yn9se3-x7t9waq74SlhLS9Ih"
	defer gock.Off() // Flush pending mocks after test execution
	// gock.Observe(gock.DumpRequest)
	gock.New("https://fcm.googleapis.com/v1/projects/auth-plus-c2b74").
		Post("/messages:send").
		Reply(200)

	gock.New("https://oauth2.googleapis.com").
		Post("/token").
		Reply(200).Body()
	provider := p.NewFirebase()
	err := provider.SendPN(DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestFirebase(t *testing.T) {
	suite.Run(t, new(FirebaseTestSuite))
}
