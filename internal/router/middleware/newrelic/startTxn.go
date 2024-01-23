package newrelic

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

const (
	NrTxnKey    = "nrTxnKey"
	NrLoggerKey = "nrLoggerKey"
)

type nrMiddleware struct {
	app *newrelic.Application
	l   *logrus.Logger
}

type NrMiddleware interface {
	StartTxn() echo.MiddlewareFunc
}

func New(l *logrus.Logger, app *newrelic.Application) NrMiddleware {
	return &nrMiddleware{app, l}
}

func (n *nrMiddleware) StartTxn() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			txn := n.app.StartTransaction(ctx.Request().URL.Path)
			defer txn.End()
			logger := n.l.WithContext(newrelic.NewContext(ctx.Request().Context(), txn))
			logger.Debugf("Starting transaction for path: %s", ctx.Request().URL.Path)

			ctx.Set(NrTxnKey, txn)
			ctx.Set(NrLoggerKey, logger)

			return next(ctx)
		}
	}
}
