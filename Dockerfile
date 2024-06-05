FROM golang:1.22.3-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go