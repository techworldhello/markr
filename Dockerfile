FROM golang:1.14

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app

RUN go install -v ./...

CMD ["go", "run", "cmd/main.go", "cmd/logger.go"]