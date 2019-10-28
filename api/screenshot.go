package api

import "github.com/labstack/echo/v4"

func registerScreenshotRoutes(g *echo.Group) {
	g.POST("/", createScreenshot)
	g.POST("/files", createScreenshot)
	g.GET("/", listScreenshots)
}

func listScreenshots(ctx echo.Context) error {
	return nil
}

func createScreenshot(ctx echo.Context) error {
	return nil
}

func createScreenshotWithFile(ctx echo.Context) error {
	return nil
}
