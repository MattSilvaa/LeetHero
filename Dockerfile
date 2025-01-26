FROM golang:1.23 AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    wget \
    chromium \
    chromium-driver \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o leethero ./cmd/leethero/main.go

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    chromium \
    chromium-driver \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/leethero /usr/local/bin/leethero

RUN chmod +x /usr/local/bin/leethero

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/leethero"]