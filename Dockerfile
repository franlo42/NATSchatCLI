#Construcción
FROM golang:1.20 as builder

WORKDIR /app
COPY nats_chat.go .
RUN GOOS=linux GOARCH=amd64 go build -o nats_chat_linux nats_chat.go
RUN GOOS=darwin GOARCH=amd64 go build -o nats_chat_mac nats_chat.go
