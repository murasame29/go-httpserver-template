package router

import (
	"html"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func (r *router) health() {
	r.engine.GET("/healthz", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}

func (r *router) info() {
	r.engine.GET("/version", func(c echo.Context) error {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return echo.ErrInternalServerError
		}

		return c.HTML(http.StatusOK, html.EscapeString(info.String()))
	})
}
