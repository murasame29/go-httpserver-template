package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type router struct {
	engine *echo.Echo
}

// NewEcho は、echo/v4 を利用した http.Handlerを返す関数です。
func NewEcho() http.Handler {
	router := &router{
		engine: echo.New(),
	}

	router.health()
	router.info()

	return router.engine
}
