#Construcción
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go get github.com/nats-io/nats.go

COPY /cmd/app/main.go .

RUN GOOS=linux GOARCH=amd64 go build -o nats_chat_linux
RUN GOOS=darwin GOARCH=amd64 go build -o nats_chat_mac
