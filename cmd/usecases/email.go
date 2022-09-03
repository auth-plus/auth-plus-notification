package usecases

import (
	m "auth-plus-notification/cmd/managers"
	"math/rand"
	"time"
)

type EmailUsecase struct {
	EmailManager *m.EmailManager
}

func NewEmailUsecase(emailManager *m.EmailManager) *EmailUsecase {
	instance := new(EmailUsecase)
	instance.EmailManager = emailManager
	return instance
}

func (e *EmailUsecase) Send(email string, content string) (bool, error) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()
	provider := e.EmailManager.ChooseProvider(random)
	return provider.SendEmail(email, content)
}
