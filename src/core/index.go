package providers

import (
	m "auth-plus-notification/core/managers"
	p "auth-plus-notification/core/providers"
	u "auth-plus-notification/core/usecases"
)

type Core struct {
	emailUsecase *u.EmailUsecase
}

func NewCore() *Core {
	sendgrid := p.NewSendgrid()
	emailManager := m.NewEmailManager(sendgrid)
	emailUsecase := u.NewEmailUsecase(emailManager)
	core := new(Core)
	core.emailUsecase = emailUsecase
	return core
}
