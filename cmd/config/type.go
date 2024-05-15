package config

import "time"

// config はアプリケーションの設定を表す構造体です。基本的には環境変数から読み込みます。
type config struct {
	Server struct {
		Host            string        `env:"HOST" envDefault:"localhost"`
		Port            int           `env:"PORT" envDefault:"8080"`
		ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
	}
}

// Config は読み込まれた設定を保持します。
var Config *config
