package services

import (
	"os"
	"time"

	"github.com/personjs/signal-demod/internal/config"
	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger(cfg config.LogConfig) {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}

	Logger = zerolog.New(output).
		Level(level).
		With().
		Timestamp().
		Logger()
}
