FROM golang:1.19-alpine as builder
RUN apk --update add build-base
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -tags netgo -a -v -o ./build/http_server ./cmd/http/http_server.go
RUN go build -tags netgo -a -v -o ./build/kafka_server ./cmd/kafka/kafka_server.go

FROM alpine:3.16.2 as deploy
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/build/http_server .
COPY --from=builder /app/build/kafka_server .
EXPOSE 5001