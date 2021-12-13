# Auth+ Notification

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=auth-plus_auth-plus-backend-notification&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=auth-plus_auth-plus-backend-notification)

This project it's a sample for notification as Email, SMS, Push Notification, Whatsapp, Telegram

## Pré-requisite

- Docker v20.10.11
- Docker Compose v1.28.4
- Go v1.17.4

## Commands

```bash

# rise/destroy all dependency
make infra/up # already create tables based on ./db/schema.sql
make infra/down # does not remove volume

# make test on the same condition where it's executed on CI
make test/ci

# developer and test enviroment
make dev

# clean
make clean/docker # prune for container, volumes and image

```

## TODO

- complete test
- Add GCP Secret Manager
- Add a GCP CLoud SQL
