.PHONY: infra/up
infra/up:
	docker-compose up -d database

.PHONY: infra/down
infra/down:
	docker-compose down

.PHONY: dev
dev:
	make infra/up
	docker-compose up -d api
	docker-compose exec api sh

.PHONY: test/ci
test/ci:
	make infra/up
	docker-compose up -d api
	docker-compose exec -T api go test ./... -v -coverpkg=./... -coverprofile=coverage.out 
	make clean/docker

.PHONY: clean/go
clean/go:
	rm go.mod
	rm go.sum
	go mod init auth-plus-notification
	go get firebase.google.com/go
	go get github.com/aws/aws-sdk-go
	go get github.com/confluentinc/confluent-kafka-go
	go get github.com/gin-contrib/cors
	go get github.com/gin-gonic/gin
	go get github.com/go-telegram-bot-api/telegram-bot-api/v5
	go get github.com/prometheus/client_golang
	go get github.com/twilio/twilio-go
	go mod tidy

.PHONY: clean/docker
clean/docker:
	make infra/down
	docker container prune -f
	docker volume prune -f
	docker image prune -f
	rm -rf db/schema.sql
	rm -f db/schema.sql