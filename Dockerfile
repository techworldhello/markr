FROM golang:1.14-alpine3.11 AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o markr .