package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var router = echo.New()

func Router() http.Handler {

	v1 := router.Group("/v1")
	ss := v1.Group("/screenshots")

	registerScreenshotRoutes(ss)

	return router
}
