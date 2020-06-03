FROM golang:1.14

# Cgo enables the creation of Go packages that call C code
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app
