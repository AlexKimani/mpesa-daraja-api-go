FROM golang:1.21.1 AS builder

LABEL authors="joealexkimani"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY config ./config
COPY src/ ./src
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mpesa-daraja-api

EXPOSE 8000

CMD ["/mpesa-daraja-api"]