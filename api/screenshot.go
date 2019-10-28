package api

import (
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/queue"
	"github.com/s4kibs4mi/snapify/repos"
	"github.com/s4kibs4mi/snapify/validators"
	"net/http"
)

func registerScreenshotRoutes(g *echo.Group) {
	g.POST("/", createScreenshot)
	g.POST("/files", createScreenshot)
	g.GET("/", listScreenshots)
}

func listScreenshots(ctx echo.Context) error {
	return nil
}

func createScreenshot(ctx echo.Context) error {
	resp := response{}

	pld, err := validators.ValidateCreateScreenshot(ctx)
	if err != nil {
		resp.Title = "Invalid data"
		resp.Status = http.StatusBadRequest
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	repo := repos.NewScreenshotRepo()

	tx := app.DB().Begin()
	data, err := repo.Create(tx, pld)
	if err != nil {
		tx.Rollback()

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	for _, x := range data {
		if err := queue.SendTakeScreenShotTask(x.ID); err != nil {
			tx.Rollback()

			resp.Title = "Failed to queue task"
			resp.Status = http.StatusInternalServerError
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}
	}

	if err := tx.Commit().Error; err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Title = "Screenshot creation is in progress"
	resp.Status = http.StatusAccepted
	resp.Data = data
	return resp.ServerJSON(ctx)
}

func createScreenshotWithFile(ctx echo.Context) error {
	return nil
}
