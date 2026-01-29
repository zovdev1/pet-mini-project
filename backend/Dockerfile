# syntax=docker/dockerfile:1.4

# ========= builder =========
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o main ./cmd


# ========= migrate =========
FROM alpine:3.20 AS migrate


RUN apk add --no-cache curl

WORKDIR /tmp

RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.20.1/migrate.linux-amd64 \
       -o /usr/local/bin/migrate \
    && chmod +x /usr/local/bin/migrate

# ========= runtime =========
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/main ./main
COPY --from=migrate /usr/local/bin/migrate /usr/local/bin/migrate
COPY .env .env


USER app

EXPOSE 8080

CMD ["./main"]