FROM golang:1.17-alpine as builder
WORKDIR /app
COPY /src .
COPY /deploy .
RUN go mod download
RUN go build -o ./build

FROM alpine:3.14 as deploy
WORKDIR /app
COPY --from=builder /app/build /app/build
EXPOSE 5000
CMD [ "./build" ]