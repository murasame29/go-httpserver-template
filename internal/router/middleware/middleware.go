package middleware

import (
	nr "github.com/murasame29/go-httpserver-template/internal/router/middleware/newrelic"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type middleware struct {
	nr.NrMiddleware
}

func New(
	l *logrus.Logger,
	nrApp *newrelic.Application,
) *middleware {
	return &middleware{
		nr.New(l, nrApp),
	}
}
