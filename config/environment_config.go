// Package config contains all configuration for application work read more on https://github.com/golang-standards/project-layout#configs
package config

import "os"

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
	Firebase FirebaseEnv
	Sendgrid SendgridEnv
	Mailgun  MailgunEnv
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
		Firebase: firebase,
		Sendgrid: sendgrid,
		Mailgun:  mailgun,
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
