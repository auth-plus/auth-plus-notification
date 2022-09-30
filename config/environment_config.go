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
	URL    string
	APIKey string
}

// SendgridEnv is environment variable for sendgrid work
type SendgridEnv struct {
	URL    string
	APIKey string
}

// FirebaseEnv is environment variable for firebase work
type FirebaseEnv struct {
	Credential string
	AppName    string
}

// ProviderEnv contains all providers configurations
type ProviderEnv struct {
	Amazon   AmazonEnv
	Firebase FirebaseEnv
	Mailgun  MailgunEnv
	Sendgrid SendgridEnv
}

// AppEnv contains all necessary property to application initiate
type AppEnv struct {
	Name string
	Port string
	Env  string
}

// PrometheusEnv for prometheus configuration
type PrometheusEnv struct {
	URL  string
	Port string
}

// Environment contains all configurations
type Environment struct {
	App        AppEnv
	Providers  ProviderEnv
	Prometheus PrometheusEnv
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
		URL:    os.Getenv("SENDGRID_URL"),
		APIKey: os.Getenv("SENDGRID_API_KEY"),
	}
	firebase := FirebaseEnv{
		Credential: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		AppName:    os.Getenv("GOOGLE_APPLICATION_NAME"),
	}
	mailgun := MailgunEnv{
		URL:    os.Getenv("MAILGUN_URL"),
		APIKey: os.Getenv("MAILGUN_API_KEY"),
	}
	providers := ProviderEnv{
		Amazon:   amazon,
		Firebase: firebase,
		Mailgun:  mailgun,
		Sendgrid: sendgrid,
	}
	//Default
	app := AppEnv{
		Name: os.Getenv("APP_NAME"),
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("GO_ENV"),
	}
	prometheus := PrometheusEnv{
		URL:  os.Getenv("PROMETHEUS_URL"),
		Port: os.Getenv("PROMETHEUS_PORT"),
	}
	//Exporting
	env := Environment{
		App:        app,
		Providers:  providers,
		Prometheus: prometheus,
	}
	return env
}
