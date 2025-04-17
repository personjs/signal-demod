# Signal Demod ✈️

A modular, Go-powered signal demodulator that reads from RTL-SDR, decodes messages, streams live via WebSocket, and optionally persists data to a database.

---

## 🚀 Features

- 📡 Real-time ADS-B decoding from RTL-SDR (1090 MHz)
- 🧠 Demodulation and message parsing in pure Go
- 🌐 Live data broadcast via WebSocket
- 🗃️ Optional database persistence with SQLite, Postgres, or MySQL
- 🐳 Docker-ready and deployable
- 🧰 CLI-driven with pluggable demod signals (`adsb`, more coming...)

---

## 📦 Installation

### Clone and build:

```bash
git clone https://github.com/personjs/signal-demod.git
cd signal-demod
go build -o signal-demod ./cmd/main
```

### Docker run:

```bash
docker run --rm \
--device /dev/bus/usb:/dev/bus/usb \
--name signal-demod \
--volume /data/signal-demod:/app/data \
--publish 8080:8080 \
-it yunostove/signal-demod
```