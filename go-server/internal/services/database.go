package services

import (
	"github.com/personjs/signal-demod/internal/config"
	"github.com/personjs/signal-demod/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cfg config.DBConfig) {
	var dialector gorm.Dialector

	switch cfg.Driver {
	case "sqlite":
		dialector = sqlite.Open(cfg.DSN)
	case "postgres":
		dialector = postgres.Open(cfg.DSN)
	case "mysql":
		dialector = mysql.Open(cfg.DSN)
	default:
		Logger.Fatal().Str("DB_DRIVER", cfg.Driver).Msg("unsupported")
	}

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		Logger.Fatal().Err(err).Msg("failed to connect to database")
	}

	// Auto-create schema
	if err := DB.AutoMigrate(&models.ADSBMessage{}); err != nil {
		Logger.Fatal().Err(err).Msg("failed to migrate database")
	}
}
