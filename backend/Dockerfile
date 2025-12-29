FROM golang:1.25.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd


RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.20.1/migrate.linux-amd64 \
    -o /usr/local/bin/schema && \
    chmod +x /usr/local/bin/schema


FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main ./main
COPY --from=builder /usr/local/bin/schema /usr/local/bin/
COPY --from=builder /app/schema ./schema
COPY .env .

EXPOSE 8080

CMD ["./main"]