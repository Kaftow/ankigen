package config

import (
	"errors"

	"github.com/spf13/viper"
)

// default values for the DB configuration.
const defaultDriver = "sqlite"
const defaultDSN = "data/app.db"

// Config describes the full application configuration surfaced to callers.
type Config struct {
	DB DBConfig `mapstructure:"db"`
}

// DBConfig configures the database driver, connection string, and pool options.
type DBConfig struct {
	Driver          string `mapstructure:"driver"`
	DSN             string `mapstructure:"dsn"`
	LogMode         bool   `mapstructure:"log_mode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// Load reads configuration from the provided path (if any), applies defaults,
// and validates the resulting configuration.
func Load(path string) (*Config, error) {
	v := viper.New()

	// Set default values
	v.SetDefault("db.driver", defaultDriver)
	v.SetDefault("db.dsn", defaultDSN)

	// Read from file if path is provided
	if path != "" {
		v.SetConfigFile(path)
		v.SetConfigType("yaml")

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// validate enforces driver-specific requirements for the database settings.
func validate(cfg *Config) error {
	switch cfg.DB.Driver {
	case "sqlite":
		if cfg.DB.DSN == "" {
			return errors.New("sqlite dsn is required")
		}

	case "mysql":
		if cfg.DB.DSN == "" {
			return errors.New("mysql dsn is required")
		}

	default:
		return errors.New("unsupported db driver: " + cfg.DB.Driver)
	}

	return nil
}
