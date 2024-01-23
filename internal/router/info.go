package router

import (
	"github.com/labstack/echo/v4"
)

func (r *router) health() {
	r.engine.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}

func (r *router) info() {
	r.engine.GET("/version", func(c echo.Context) error {
		return c.String(200, r.version)
	})
}
