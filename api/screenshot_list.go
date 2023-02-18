package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s4kibs4mi/snapify/apimodels"
	"net/http"
	"time"
)

func (h *handlers) ScreenshotList(ctx *fiber.Ctx) error {
	query := struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}{}
	if err := ctx.QueryParser(&query); err != nil {
		return h.Serve(ctx, http.StatusBadRequest, map[string]interface{}{"err": err})
	}

	screenshots, err := h.screenshotDao.List(query.Page, query.Limit)
	if err != nil {
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	var formattedScreenshots []apimodels.RespScreenshot
	for _, s := range screenshots {
		formattedScreenshots = append(formattedScreenshots, apimodels.RespScreenshot{
			ID:        s.ID.String(),
			URL:       s.URL,
			Status:    string(s.Status),
			CreatedAt: s.CreatedAt.Format(time.RFC3339),
		})
	}

	return h.Serve(ctx, http.StatusOK, fiber.Map{
		"data": formattedScreenshots,
	})
}
