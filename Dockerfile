FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o golos ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/golos .
COPY --from=builder /app/config.yaml .   
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env .env

EXPOSE 8080

CMD ["./golos"]