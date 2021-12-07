FROM golang:1.17-alpine

WORKDIR /app

COPY /src .

RUN go mod download

RUN go build -o /build

EXPOSE 5000

CMD [ "./build" ]