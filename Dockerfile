FROM golang:1.21.1 AS builder

LABEL authors="joealexkimani"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.0

COPY src/ ./src
COPY *.go ./

RUN swag init --output /docs/mpesa

RUN CGO_ENABLED=0 GOOS=linux go build -o /mpesa-daraja-api

FROM scratch

COPY --from=builder /mpesa-daraja-api /mpesa-daraja-api

WORKDIR /app
COPY --from=builder /docs/mpesa /docs/mpesa
COPY migrations ./migrations
COPY config ./config

EXPOSE 8080

CMD ["/mpesa-daraja-api"]