package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s4kibs4mi/snapify/apimodels"
	"net/http"
	"time"
)

// ScreenshotList is a function to list screenshots
// @Summary Retrieve screenshots
// @Description Retrieve screenshots
// @Param	Token	header	string	true	"Authentication header"
// @Param	limit	query	string	false	"Number of items"
// @Param	page	query	string	false	"Page index"
// @Tags screenshots
// @Produce json
// @Success 200 {object} apimodels.RespScreenshotList{data=[]apimodels.RespScreenshotData}
// @Router /v1/screenshots [get]
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

	var formattedScreenshots []apimodels.RespScreenshotData
	for _, s := range screenshots {
		formattedScreenshots = append(formattedScreenshots, apimodels.RespScreenshotData{
			ID:        s.ID.String(),
			URL:       s.URL,
			Status:    string(s.Status),
			CreatedAt: s.CreatedAt.Format(time.RFC3339),
		})
	}

	return h.Serve(ctx, http.StatusOK, apimodels.RespScreenshotList{
		Data: formattedScreenshots,
	})
}
