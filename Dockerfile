FROM golang:1.21.1 AS builder

LABEL authors="joealexkimani"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY src/ ./src
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mpesa-daraja-api

FROM scratch

COPY --from=builder /mpesa-daraja-api /mpesa-daraja-api

WORKDIR /app
COPY migrations ./migrations
COPY config ./config

EXPOSE 8000

CMD ["/mpesa-daraja-api"]