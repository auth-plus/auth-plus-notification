# Auth+ Notification

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=auth-plus_auth-plus-backend-notification&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=auth-plus_auth-plus-backend-notification)

This project it's a sample for notification as Email, SMS, Push Notification, Whatsapp and Telegram.

## Pré-requisite

- Docker v20.10.11
- Docker Compose v1.28.4
- Go v1.19

## Enviroment Variables

Please follow example.env

## Commands with Docker Setup

```bash
# make test on the same condition where it's executed on CI
make test/ci

# developer on docker
make dev

# prune for container, volumes and image
make clean/docker 
```

## Commands with Local Setup

```bash
# install dependecies on local
go mod download

# run server on local
go run ./server.go
```

## How to update packages

```bash
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
