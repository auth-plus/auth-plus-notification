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
	firebase := p.NewFirebase()
	mailgun := p.NewMailgun()
	sendgrid := p.NewSendgrid()
	sns := p.NewSNS()
	telegram := p.NewTelegram()
	whatsapp := p.NewWhatsapp()

	//Managers
	emailManager := m.NewEmailManager(sendgrid, mailgun)
	pushNotificationManager := m.NewPushNotificationManager(firebase)
	smsManager := m.NewSmsManager(sns)
	telegramManager := m.NewTelegramManager(telegram)
	whatsappManager := m.NewWhatsappManager(whatsapp)

	//Usecases
	emailUsecase := u.NewEmailUsecase(emailManager)
	pushNotificationUsecase := u.NewPushNotificationUsecase(pushNotificationManager)
	smsUsecase := u.NewSmsUsecase(smsManager)
	telegramUsecase := u.NewTelegramUsecase(telegramManager)
	whatsappUsecase := u.NewWhatsappUsecase(whatsappManager)

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
