FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY config ./config

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/config /app/config

CMD ["./app"]
