package config

type LogConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"info"` // debug, info, warn, error
}
