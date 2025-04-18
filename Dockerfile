# Stage 1: Build npm app
FROM node:lts-alpine as builder

WORKDIR /app
COPY frontend/ ./
RUN npm install && npm run build

# Stage 2: Serve with Go
FROM golang:1.24-bookworm AS server

RUN apt-get update \
&& apt-get install -y git build-essential pkg-config librtlsdr-dev libusb-1.0-0-dev

WORKDIR /app
COPY go-server/ .
COPY --from=builder /app/dist ./public

RUN go build -o server ./cmd/main

EXPOSE 8080
EXPOSE 8081

CMD ["./server", "run", "adsb"]