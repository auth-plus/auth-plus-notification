// Package config contains all configuration for application work read more on https://github.com/golang-standards/project-layout#configs
package config

import "os"

// AmazonEnv is environment variable for sns work
type AmazonEnv struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// MailgunEnv is environment variable for mailgun work
type MailgunEnv struct {
	APIKey string
}

// OnesignalEnv is environment variable for onesignal work
type OnesignalEnv struct {
	APIKey string
}

// SendgridEnv is environment variable for sendgrid work
type SendgridEnv struct {
	APIKey string
}

// FirebaseEnv is environment variable for firebase work
type FirebaseEnv struct {
	Credential string
	AppName    string
}

// TelegramEnv is environment variable for mailgun work
type TelegramEnv struct {
	APIKey string
}

// TwilioEnv is environment variable for mailgun work
type TwilioEnv struct {
	AccountID string
	Token     string
}

// ProviderEnv contains all providers configurations
type ProviderEnv struct {
	Amazon    AmazonEnv
	Firebase  FirebaseEnv
	Mailgun   MailgunEnv
	Onesignal OnesignalEnv
	Sendgrid  SendgridEnv
	Telegram  TelegramEnv
	Twilio    TwilioEnv
}

// AppEnv contains all necessary property to application initiate
type AppEnv struct {
	Name string
	Port string
	Env  string
}

// KafkaEnv for kafka configuration
type KafkaEnv struct {
	URL  string
	Port string
}

// Environment contains all configurations
type Environment struct {
	App       AppEnv
	Providers ProviderEnv
	Kafka     KafkaEnv
}

// GetEnv exports env config instead of multiplaces to maintain
func GetEnv() Environment {

	// Provider
	amazon := AmazonEnv{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
	}
	sendgrid := SendgridEnv{
		APIKey: os.Getenv("SENDGRID_API_KEY"),
	}
	firebase := FirebaseEnv{
		Credential: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		AppName:    os.Getenv("GOOGLE_APPLICATION_NAME"),
	}
	mailgun := MailgunEnv{
		APIKey: os.Getenv("MAILGUN_API_KEY"),
	}
	onesignal := OnesignalEnv{
		APIKey: os.Getenv("ONESIGNAL_API_KEY"),
	}
	telegram := TelegramEnv{
		APIKey: os.Getenv("TELEGRAM_API_KEY"),
	}
	twilio := TwilioEnv{
		AccountID: os.Getenv("TWILIO_ACCOUNT_SID"),
		Token:     os.Getenv("TWILIO_AUTH_TOKEN"),
	}
	providers := ProviderEnv{
		Amazon:    amazon,
		Firebase:  firebase,
		Mailgun:   mailgun,
		Onesignal: onesignal,
		Sendgrid:  sendgrid,
		Telegram:  telegram,
		Twilio:    twilio,
	}
	//Default
	app := AppEnv{
		Name: os.Getenv("APP_NAME"),
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("GO_ENV"),
	}
	kafka := KafkaEnv{
		URL:  os.Getenv("KAFKA_URL"),
		Port: os.Getenv("KAFKA_PORT"),
	}
	//Exporting
	env := Environment{
		App:       app,
		Providers: providers,
		Kafka:     kafka,
	}
	return env
}
