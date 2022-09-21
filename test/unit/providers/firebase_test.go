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
	err := faker.FakeData(&mockData)
	if err != nil {
		fmt.Println(err)
	}

	const DeviceID string = "eqEiHBNFOisH1w2wlmZgUH:APA91bGV0clbAlzqgzgRJOFXawN1Kp4163nwjRSWZ46pNXnfBbwEiaV-Bop4R2YrnliXPMic3XXa9N7DU_orVMPjc7bPrgJcGk9bqwRzJYfR62J5fGdg_umaeSyHTyyTQeDGhRoBSsjR"
	defer gock.Off() // Flush pending mocks after test execution
	gock.Observe(gock.DumpRequest)
	gock.New("https://fcm.googleapis.com/v1/projects/auth-plus-c2b74").
		Post("/messages:send").
		Reply(200)

	provider := p.NewFirebase()
	resp, err := provider.SendPN(DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), resp, true)
	assert.Equal(suite.T(), err, nil)
}

func TestFirebase(t *testing.T) {
	suite.Run(t, new(FirebaseTestSuite))
}
