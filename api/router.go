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

	router.GET("", func(ctx echo.Context) error {
		return ctx.HTML(http.StatusOK, "<h1>Ok</h1>")
	})

	v1 := router.Group("/v1")
	ss := v1.Group("/screenshots")

	registerScreenshotRoutes(ss)

	return router
}
