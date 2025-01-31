FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o leethero ./cmd/leethero/main.go

FROM chromedp/headless-shell:latest

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/leethero /usr/local/bin/leethero
RUN chmod +x /usr/local/bin/leethero

ENV CHROME_FLAGS="\
  --no-sandbox \
  --disable-dev-shm-usage \
  --disable-gpu \
  --disable-software-rasterizer \
  --disable-extensions \
  --disable-default-apps \
  --disable-translate \
  --disable-sync \
  --disable-background-networking \
  --safebrowsing-disable-auto-update \
  --disable-client-side-phishing-detection \
  --mute-audio \
  --no-first-run \
  --headless"

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/leethero"]