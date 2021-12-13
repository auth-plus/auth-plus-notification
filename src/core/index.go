package providers

import (
	m "auth-plus-notification/core/managers"
	p "auth-plus-notification/core/providers"
	u "auth-plus-notification/core/usecases"
)

type Core struct {
	EmailUsecase            *u.EmailUsecase
	PushNotificationUsecase *u.PushNotificationUsecase
	SmsUsecase              *u.SmsUsecase
	TelegramUsecase         *u.TelegramUsecase
	WhatsappUsecase         *u.WhatsappUsecase
}

func NewCore() Core {
	//Providers
	sendgrid := p.NewSendgrid()
	mailgun := p.NewMailgun()

	//Managers
	emailManager := m.NewEmailManager(sendgrid, mailgun)

	//Usecases
	emailUsecase := u.NewEmailUsecase(emailManager)
	pushNotificationUsecase := u.NewPushNotificationUsecase()
	smsUsecase := u.NewSmsUsecase()
	telegramUsecase := u.NewTelegramUsecase()
	whatsappUsecase := u.NewWhatsappUsecase()

	//Constructing Core
	core := Core{
		EmailUsecase:            emailUsecase,
		PushNotificationUsecase: pushNotificationUsecase,
		SmsUsecase:              smsUsecase,
		TelegramUsecase:         telegramUsecase,
		WhatsappUsecase:         whatsappUsecase,
	}
	return core
}
