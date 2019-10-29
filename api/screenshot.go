package api

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/models"
	"github.com/s4kibs4mi/snapify/queue"
	"github.com/s4kibs4mi/snapify/repos"
	"github.com/s4kibs4mi/snapify/services"
	"github.com/s4kibs4mi/snapify/validators"
	"net/http"
	"strconv"
)

func registerScreenshotRoutes(g *echo.Group) {
	g.POST("/", createScreenshot)
	g.POST("/files/", createScreenshotWithFile)
	g.GET("/", listScreenshots)
	g.GET("/:id/", getScreenshot)
}

func listScreenshots(ctx echo.Context) error {
	resp := response{}

	limitQ := ctx.Request().URL.Query().Get("limit")
	pageQ := ctx.Request().URL.Query().Get("page")

	limit, err := strconv.ParseInt(limitQ, 10, 32)
	if err != nil {
		limit = 10
	}
	page, err := strconv.ParseInt(pageQ, 10, 32)
	if err != nil {
		page = 1
	}

	repo := repos.NewScreenshotRepo()
	db := app.DB()

	total, err := repo.Count(db)
	if err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	totalPages := int64(total) / limit
	if int64(total)%limit != 0 {
		totalPages++
	}

	data, err := repo.List(db, page, limit)
	if err != nil {
		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = map[string]interface{}{
		"screenshots": data,
		"meta": map[string]interface{}{
			"total":        total,
			"current_page": page,
			"page_limit":   limit,
			"total_pages":  totalPages,
		},
	}
	return resp.ServerJSON(ctx)
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
	resp := response{}

	pld, err := validators.ValidateCreateScreenshotFromFile(ctx)

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

func getScreenshot(ctx echo.Context) error {
	resp := response{}

	repo := repos.NewScreenshotRepo()
	db := app.DB()

	ID := ctx.Param("id")
	m, err := repo.Get(db, ID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Title = "Screenshot not found"
			resp.Status = http.StatusNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	if m.Status == models.Queued || m.Status == models.Failed {
		resp.Title = "Invalid screenshot"
		resp.Status = http.StatusBadRequest
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	o, err := services.ServeAsStreamFromMinio(m.StoredPath)
	if err != nil {
		resp.Title = "Minio service failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	return resp.ServerImageFromMinio(ctx, o)
}
