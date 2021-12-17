package managers

//ENUM for Email
type EnumEmailProvider string

const (
	SendGrid EnumEmailProvider = "SendGrid"
	Mailgun  EnumEmailProvider = "Mailgun"
)

//ENUM for PushNotification
type EnumPushNotificationProvider string

const (
	Firebase EnumPushNotificationProvider = "Firebase"
	Braze    EnumPushNotificationProvider = "Braze"
)

//ENUM for Sms
type EnumSmsProvider string

const (
	SNS EnumSmsProvider = "Sns"
)

//ENUM for Telegram
type EnumTelegramProvider string

const (
	Telegram EnumTelegramProvider = "Telegram"
)

//ENUM for Whatsapp
type EnumWhatsappProvider string

const (
	Whatsapp EnumWhatsappProvider = "Whatsapp"
)
