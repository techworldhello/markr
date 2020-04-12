FROM golang:1.14-alpine3.11 AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o markr .
