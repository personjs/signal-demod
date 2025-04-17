package config

type DBConfig struct {
	Driver  string `env:"DB_DRIVER" envDefault:"sqlite"`       // sqlite, postgres, mysql
	DSN     string `env:"DB_DSN" envDefault:"signal-demod.db"` // path or connection string
	MaxOpen int    `env:"DB_MAX_OPEN" envDefault:"10"`
}
