package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadEnv(path ...string) error {
	if len(path) > 0 {
		if err := godotenv.Load(path...); err != nil {
			return fmt.Errorf("failed to load env: %w", err)
		}
	}

	config := &config{}

	if err := env.Parse(&config.Server); err != nil {
		return err
	}

	// ... configに構造体を追加したら env.Parseする
	Config = config

	return nil
}
