package db

import (
	"os"
	"path/filepath"
	"time"

	"ankigen/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newSQLite(cfg config.DBConfig) (*gorm.DB, error) {
	// Auto create directory if not exists
	if err := os.MkdirAll(filepath.Dir(cfg.DSN), 0755); err != nil {
		return nil, err
	}

	gormCfg := &gorm.Config{
		PrepareStmt: true,
	}
	if cfg.LogMode {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(sqlite.Open(cfg.DSN), gormCfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxOpen := cfg.MaxOpenConns
	if maxOpen <= 0 {
		maxOpen = 1
	}
	sqlDB.SetMaxOpenConns(maxOpen)

	maxIdle := cfg.MaxIdleConns
	if maxIdle <= 0 {
		maxIdle = 1
	}
	sqlDB.SetMaxIdleConns(maxIdle)

	if cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	}

	db.Exec("PRAGMA journal_mode = WAL;")
	db.Exec("PRAGMA foreign_keys = ON;")

	return db, nil
}
