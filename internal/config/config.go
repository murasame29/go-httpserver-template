package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config はアプリケーションの設定を表す構造体です。
type Config struct {
	Server ServerConfig
}

// ServerConfig はサーバーの設定を表す構造体です。
type ServerConfig struct {
	Host            string        `env:"HOST" envDefault:"localhost"`
	Port            int           `env:"PORT" envDefault:"8080"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" envDefault:"30s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" envDefault:"30s"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" envDefault:"60s"`
}

// Load は環境変数から設定を読み込みます。
func Load(paths ...string) (*Config, error) {
	if len(paths) > 0 {
		if err := godotenv.Load(paths...); err != nil {
			return nil, fmt.Errorf("failed to load env: %w", err)
		}
	}

	cfg := &Config{}

	if err := env.Parse(&cfg.Server); err != nil {
		return nil, fmt.Errorf("failed to parse server config: %w", err)
	}

	return cfg, nil
}
