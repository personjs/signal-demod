FROM golang:1.24-bookworm AS builder

RUN apt-get update \
&& apt-get install -y git build-essential pkg-config librtlsdr-dev libusb-1.0-0-dev

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o signal-demod ./cmd/main

FROM debian:bookworm-slim

# Install minimal dependencies
RUN apt-get update \
&& apt-get install -y librtlsdr0 libusb-1.0-0 ca-certificates tzdata \
&& rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 👤 Create non-root user
RUN useradd --system --create-home --shell /usr/sbin/nologin demod

# Copy binary and default data location
COPY --from=builder /app/signal-demod .

# Ensure data dir exists for SQLite
RUN mkdir -p /app/data

# ✅ Set ownership to demod user
RUN chown -R demod:demod /app

ENV DB_DSN=data/signal-demod.db

EXPOSE 8080

CMD ["./signal-demod", "run", "adsb"]
    