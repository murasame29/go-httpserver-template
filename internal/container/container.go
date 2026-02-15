package container

import (
	"net/http"

	"github.com/murasame29/go-httpserver-template/internal/config"
	"github.com/murasame29/go-httpserver-template/internal/router"
	"github.com/murasame29/go-httpserver-template/internal/server"
	"go.uber.org/dig"
)

var container *dig.Container

type provideArg struct {
	constructor any
	opts        []dig.ProvideOption
}

// New は dig を用いて依存注入コンテナを初期化します。
func New(cfg *config.Config) error {
	container = dig.New()

	args := []provideArg{
		// Config
		{constructor: func() *config.Config { return cfg }},
		{constructor: func() *config.ServerConfig { return &cfg.Server }},

		// Router
		{constructor: router.NewServeMux, opts: as[http.Handler]()},

		// Server
		{constructor: server.New},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
}

func as[T any]() []dig.ProvideOption {
	return []dig.ProvideOption{dig.As(new(T))}
}

// Invoke は *dig.Container の Invoke をラップした関数です。
func Invoke[T any](opts ...dig.InvokeOption) (T, error) {
	var r T
	if err := container.Invoke(func(t T) error {
		r = t
		return nil
	}, opts...); err != nil {
		return r, err
	}
	return r, nil
}
