FROM golang:1.19-alpine as dependency
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod vendor

FROM golang:1.19-alpine as builder
WORKDIR /app
RUN apk --update add build-base
COPY --from=dependency /app/vendor /app/vendor
COPY . .
RUN go build -o ./build/server server.go

FROM alpine:3.16.2 as deploy
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/build/server .
EXPOSE 5001
CMD [ "./server" ]