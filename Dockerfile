FROM golang:1.23-alpine AS builder
RUN apk --update add build-base
WORKDIR /app
COPY ./api ./api
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./test ./test
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum 
RUN go mod download
RUN go build -tags netgo -a -v -o ./build/http_server ./cmd/http/http_server.go
RUN go build -tags netgo -a -v -o ./build/kafka_server ./cmd/kafka/kafka_server.go

FROM alpine:3.16.2 AS deploy
RUN apk --no-cache add ca-certificates
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
WORKDIR /app
COPY --from=builder /app/build/http_server .
COPY --from=builder /app/build/kafka_server .
USER nonroot
EXPOSE 5001