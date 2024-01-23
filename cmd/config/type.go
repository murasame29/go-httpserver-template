package config

import "time"

// config はアプリケーションの設定を表す構造体です。基本的には環境変数から読み込みます。
type config struct {
	Server struct {
		Host            string        `envconfig:"HOST" default:"localhost"`
		Port            int           `envconfig:"PORT" default:"8080"`
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`
	}

	NewRelic struct {
		AppName string `envconfig:"NEW_RELIC_APP_NAME" default:"golang-backend-server"`
		License string `envconfig:"NEW_RELIC_LICENSE" default:""`
	}
}

// Config は読み込まれた設定を保持します。
var Config *config
