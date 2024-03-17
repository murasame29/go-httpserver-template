package config

import "time"

// config はアプリケーションの設定を表す構造体です。基本的には環境変数から読み込みます。
type config struct {
	Server struct {
		Host            string        `env:"HOST" envDefault:"localhost"`
		Port            int           `env:"PORT" envDefault:"8080"`
		ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
	}

	NewRelic struct {
		AppName string `env:"NEW_RELIC_APP_NAME" envDefault:"golang-backend-server"`
		License string `env:"NEW_RELIC_LICENSE" envDefault:""`
	}
}

// Config は読み込まれた設定を保持します。
var Config *config
