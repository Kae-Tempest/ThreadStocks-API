FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/main .
# COPY --from=builder /app/.env .env

EXPOSE 8080

CMD ["./main"]