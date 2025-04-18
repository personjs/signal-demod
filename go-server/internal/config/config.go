package config

import (
	"github.com/caarlos0/env/v10"
)

var App Config

type Config struct {
	Log LogConfig
	DB  DBConfig
}

func Load() error {
	App = Config{}

	if err := env.Parse(&App.Log); err != nil {
		return err
	}
	if err := env.Parse(&App.DB); err != nil {
		return err
	}

	return nil
}
