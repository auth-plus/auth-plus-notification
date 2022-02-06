package config

import "os"

type MailgunEnv struct {
	Url    string
	ApiKey string
}
type SendgridEnv struct {
	Url    string
	ApiKey string
}
type FirebaseEnv struct {
	CredentialPath string
	AppName        string
}

type ProviderEnv struct {
	Firebase FirebaseEnv
	Sendgrid SendgridEnv
	Mailgun  MailgunEnv
}

type AppEnv struct {
	Name string
	Port string
	Env  string
}

type PrometheusEnv struct {
	Url  string
	Port string
}

type Environment struct {
	App        AppEnv
	Providers  ProviderEnv
	Prometheus PrometheusEnv
}

func GetEnv() Environment {
	app := AppEnv{
		Name: os.Getenv("APP_NAME"),
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("GO_ENV"),
	}
	prometheus := PrometheusEnv{
		Url:  os.Getenv("PROMETHEUS_URL"),
		Port: os.Getenv("PROMETHEUS_PORT"),
	}
	sendgrid := SendgridEnv{
		Url:    os.Getenv("SENDGRID_URL"),
		ApiKey: os.Getenv("SENDGRID_API_KEY"),
	}
	firebase := FirebaseEnv{
		CredentialPath: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		AppName:        os.Getenv("GOOGLE_APPLICATION_NAME"),
	}
	mailgun := MailgunEnv{
		Url:    os.Getenv("MAILGUN_URL"),
		ApiKey: os.Getenv("MAILGUN_API_KEY"),
	}
	providers := ProviderEnv{
		Firebase: firebase,
		Sendgrid: sendgrid,
		Mailgun:  mailgun,
	}
	env := Environment{
		App:        app,
		Providers:  providers,
		Prometheus: prometheus,
	}
	return env
}
