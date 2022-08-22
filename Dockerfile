FROM golang:1.17-alpine as dependency
WORKDIR /app
COPY /src .
COPY /deploy .
RUN go mod download

FROM dependency as builder
RUN go build -o ./server

FROM alpine:3.14 as deploy
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 5000
CMD [ "./server" ]