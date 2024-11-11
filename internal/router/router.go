package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewEcho は、echo/v4 を利用した http.Handlerを返す関数です。
func NewEcho() http.Handler {
	engine := echo.New()

	engine.GET("/healthz", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	return engine
}
