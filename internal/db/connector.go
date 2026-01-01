package db

import (
	"fmt"

	"ankigen/internal/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg config.DBConfig) error {
	var (
		database *gorm.DB
		err      error
	)

	switch cfg.Driver {
	case "", "sqlite":
		database, err = newSQLite(cfg)

	case "mysql":
		database, err = newMySQL(cfg)

	default:
		return fmt.Errorf("unsupported db driver: %s", cfg.Driver)
	}

	if err != nil {
		return err
	}

	DB = database
	return nil
}
