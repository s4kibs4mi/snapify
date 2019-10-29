package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var router = echo.New()

func Router() http.Handler {
	router.Pre(middleware.AddTrailingSlash())
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	v1 := router.Group("/v1")
	ss := v1.Group("/screenshots")

	registerScreenshotRoutes(ss)

	return router
}
