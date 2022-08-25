# Auth+ Notification

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=auth-plus_auth-plus-backend-notification&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=auth-plus_auth-plus-backend-notification)

This project it's a sample for notification as Email, SMS, Push Notification, Whatsapp, Telegram

## Pré-requisite

- Docker v20.10.11
- Docker Compose v1.28.4
- Go v1.17.4

## Enviroment Variables

Please follow example.env

## Commands with Docker Setup

```bash

# make test on the same condition where it's executed on CI
make test/ci

# developer and test enviroment on docker
make dev

# install dependecies on local
go mod download

# run server on local
go run ./server.go

# clean
make clean/docker # prune for container, volumes and image

```

## How to update packages

```bash

go get -u firebase.google.com/go v3.13.0+incompatible
go get -u github.com/confluentinc/confluent-kafka-go v1.9.2
go get -u github.com/gin-contrib/cors v1.4.0
go get -u github.com/gin-gonic/gin v1.8.1
go get -u github.com/prometheus/client_golang
go get -u golang.org/x/lint/golint
go mod tidy
```

## TODO

- complete test
- Add GCP Secret Manager
- Add a GCP CLoud SQL
