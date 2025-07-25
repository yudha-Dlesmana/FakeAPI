FROM golang:1.24.5-alpine3.21

WORKDIR /usr/src/app

COPY go.mod go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/app ./...