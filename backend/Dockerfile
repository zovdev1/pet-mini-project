FROM golang:1.25.1-alpine AS builder

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

RUN apk add --no-cache git curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd


RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.20.1/migrate.linux-amd64 \
    -o /usr/local/bin/schema && \
    chmod +x /usr/local/bin/schema

RUN chown -R app:app /app


# ========= runtime =========
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/main ./main
COPY --from=builder /usr/local/bin/schema /usr/local/bin/
COPY --from=builder /app/schema ./schema
COPY .env .env

RUN chown -R app:app /app /usr/local/bin/schema

USER app

EXPOSE 8080

CMD ["./main"]