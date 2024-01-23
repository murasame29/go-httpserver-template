package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

const DefaultVersion = "version not set yet"

type router struct {
	version string

	l      *logrus.Logger
	nrApp  *newrelic.Application
	engine *echo.Echo
}

// Option は、router の設定を行うための関数です。
type Option func(*router)

// WithVersion　は、router のバージョンを設定するための関数です。
func WithVersion(version string) Option {
	return func(r *router) {
		r.version = version
	}
}

// WithLogger は、router のロガーを設定するための関数です。
func WithLogger(l *logrus.Logger) Option {
	return func(r *router) {
		r.l = l
	}
}

// WithNewRelic は、router の NewRelic を設定するための関数です。
func WithNewRelic(nrApp *newrelic.Application) Option {
	return func(r *router) {
		r.nrApp = nrApp
	}
}

// NewEchoServer は、echo/v4 を利用した http.Handlerを返す関数です。
func NewEchoServer(opts ...Option) http.Handler {
	router := &router{
		version: DefaultVersion,
		engine:  echo.New(),
	}

	for _, opt := range opts {
		opt(router)
	}

	// if router.nrApp != nil {
	// 	router.engine.Use(
	// 		middleware.New(router.l, router.nrApp).StartTxn(),
	// 		nrecho.Middleware(router.nrApp),
	// 	)
	// }

	router.health()
	router.info()

	return router.engine
}
