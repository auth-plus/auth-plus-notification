// Package cmd read on https://github.com/golang-standards/project-layout#cmd
package cmd

import (
	m "auth-plus-notification/cmd/managers"
	p "auth-plus-notification/cmd/providers"
	u "auth-plus-notification/cmd/usecases"
)

// Core contains all usecases
type Core struct {
	EmailUsecase            *u.EmailUsecase
	PushNotificationUsecase *u.PushNotificationUsecase
	SmsUsecase              *u.SmsUsecase
	TelegramUsecase         *u.TelegramUsecase
	WhatsappUsecase         *u.WhatsappUsecase
}

// NewCore is a function that make middle between presentation layer and the usecases
func NewCore() Core {
	//Providers
	firebase := p.NewFirebase()
	mailgun := p.NewMailgun()
	sendgrid := p.NewSendgrid()
	sns := p.NewSNS()
	telegram := p.NewTelegram()
	twilio := p.NewTwilio()
	onesignal := p.NewOneSignal()

	//Managers
	// emailManager := m.NewRandomEmailManager(sendgrid, mailgun, onesignal)
	emailManager := m.NewIPWarmimgmailManager(sendgrid, mailgun, onesignal)
	pushNotificationManager := m.NewRandomPushNotificationManager(firebase, onesignal)
	smsManager := m.NewRandomSmsManager(sns, onesignal)
	telegramManager := m.NewRandomTelegramManager(telegram)
	whatsappManager := m.NewRandomWhatsappManager(twilio)

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
