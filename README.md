# Auth+ Notification

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=auth-plus_auth-plus-backend-notification&metric=coverage)](https://sonarcloud.io/summary/new_code?id=auth-plus_auth-plus-backend-notification)

[![Test Coverage](https://api.codeclimate.com/v1/badges/8b06e8bee2391dc8817a/test_coverage)](https://codeclimate.com/github/auth-plus/auth-plus-backend-notification/test_coverage)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/870535e320a4452eac49e677bd5025de)](https://www.codacy.com/gh/auth-plus/auth-plus-backend-notification/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=auth-plus/auth-plus-backend-notification&amp;utm_campaign=Badge_Grade)

This project it's a sample for notification as Email, SMS, Push Notification, Whatsapp and Telegram.

## Pré-requisite

- Docker v20.10.11
- Docker Compose v1.28.4
- Go v1.19

## Enviroment Variables

Please follow example.env

## Commands with Docker Setup

```makefile
# make test on the same condition where it's executed on CI
make test/ci

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
go run ./server.go

# run test
go test ./... -v -coverpkg=./... -coverprofile=coverage.out 
# take a look on coverage file in html after test
go tool cover -html=coverage.out -o cover.html
```

## How to update packages

```go
go get -u
go mod tidy
```

## How reset module / packages

```bash
make clean/go
```

## TODO

- Complete Test
- Add Template
