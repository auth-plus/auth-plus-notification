package test

import (
	p "auth-plus-notification/internal/providers"
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
	const AccesToken string = "ya29.c.b0AUFJQsHw3El5fuuJm-cbWeovCgNbnLlP0hW_OpRTsM1P-7bZHdeVIhAoLXvvLEDdaiXZxAuaDyWmhEQS9ENL-vF5zNs-fhUx7rHV9fnJUgm3PtyTTC3jEHun5EAIpyZgMPf2JRfvHVIcebsRMSLWT5NN3sAlLU1jIIa1anQQH8hJKaWLR1w7YPdbnGUlikk7v2YZN0jGz5MiYEPTv3QAHUUwrBe6aXA"
	const DeviceID string = "cDxDrGiXRnMXFXFsyDLSY5:APA91bF9QI6-YU2eJr7JcF8u6lrAGAIBEpG4j3IOrU2h2EFUahYM1z0fMR_IyqybrOfc62ASy5uRyg1uzjR2trlpmZujQ79-QQAf7iSYj_4HZn3fWYg9yn9se3-x7t9waq74SlhLS9Ih"

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://fcm.googleapis.com/v1/projects/auth-plus-c2b74").
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", AccesToken)).
		Post("/messages:send").
		Reply(200)

	gock.New("https://oauth2.googleapis.com").
		Post("/token").
		Reply(200).
		JSON(map[string]interface{}{
			"access_token": AccesToken,
			"expires_in":   3599,
			"token_type":   "Bearer",
		})
	provider := p.NewFirebase()
	err := provider.SendPN(DeviceID, mockData.Title, mockData.Content)
	assert.Equal(suite.T(), err, nil)
}

func TestFirebase(t *testing.T) {
	suite.Run(t, new(FirebaseTestSuite))
}
