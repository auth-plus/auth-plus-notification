# Auth+ Notification

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=auth-plus_auth-plus-notification&metric=coverage)](https://sonarcloud.io/summary/new_code?id=auth-plus_auth-plus-notification)

[![Test Coverage](https://api.codeclimate.com/v1/badges/7747782d29adc97edda2/test_coverage)](https://codeclimate.com/github/auth-plus/auth-plus-notification/test_coverage)

[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/870535e320a4452eac49e677bd5025de)](https://www.codacy.com/gh/auth-plus/auth-plus-notification/dashboard?utm_source=github.com&utm_medium=referral&utm_content=auth-plus/auth-plus-notification&utm_campaign=Badge_Coverage)

This project it's a sample for notification system.
In this application you cand send:

- Email (Sendgrid, Mailgun, Onesignal)
- SMS (Amazon SNS, Onesignal)
- Push Notification (Firebase, Onesignal)
- Whatsapp (Twilio)
- Telegram (Telegram API)

## Pré-requisite

- Docker v20.10.11
- Docker Compose v1.28.4
- Go v1.19

## Enviroment Variables

Please follow example.env

## Commands with Docker Setup

```makefile
# make test on the same condition where it's executed on CI
make test

# developer on docker
make dev

# prune for container, volumes and image
make clean/docker 
```

## Commands with Local Setup

```go
# install dependecies on local
go mod download

# run server on local
go run ./cmd/http_server.go

# run test
go test ./... -v -coverpkg=./... -coverprofile=c.out

# take a look on coverage file in html after test
go tool cover -html=c.out -o cover.html

# run lint
$HOME/go/bin/revive -formatter friendly ./...
```

## How to update packages

```bash
go get -u
go mod tidy
```

## TODO

- Add Template
