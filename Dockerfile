FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
