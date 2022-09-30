FROM golang:1.19-alpine as builder
RUN apk --update add build-base
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./build/server ./cmd/http_server.go

FROM alpine:3.16.2 as deploy
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/build/server .
EXPOSE 5001
CMD [ "./server" ]