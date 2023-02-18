package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/apimodels"
	"github.com/s4kibs4mi/snapify/ent"
	"net/http"
	"time"
)

// ScreenshotGet is a function to retrieve screenshot
// @Summary Retrieve screenshot info
// @Description Retrieve screenshot info
// @Param	Token	header	string	true	"Authentication header"
// @Param	screenshot_id	path	string	true	"Screenshot UUID"
// @Tags screenshots
// @Produce json
// @Success 200 {object} apimodels.RespScreenshot{data=apimodels.RespScreenshotData}
// @Router /v1/screenshots/{screenshot_id} [get]
func (h *handlers) ScreenshotGet(ctx *fiber.Ctx) error {
	screenshotID := ctx.Params("id", "")
	screenshotUUID, err := uuid.Parse(screenshotID)
	if err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusBadRequest, map[string]interface{}{"err": err})
	}

	ss, err := h.screenshotDao.Get(screenshotUUID)
	if err != nil {
		h.logger.Error(err)

		if ent.IsNotFound(err) {
			return h.Serve(ctx, http.StatusNotFound, map[string]interface{}{"err": err})
		}
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	if ss.StoredPath == nil {
		return h.Serve(ctx, http.StatusBadRequest, fiber.Map{
			"err": "stored_path is nil",
		})
	}

	signedUrl, err := h.storageService.FetchUrl(*ss.StoredPath)
	if err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	h.logger.Info("SignedUrl: ", signedUrl)

	return h.Serve(ctx, http.StatusOK, apimodels.RespScreenshot{
		Data: apimodels.RespScreenshotData{
			ID:            ss.ID.String(),
			URL:           ss.URL,
			Status:        string(ss.Status),
			ScreenshotURL: &signedUrl,
			CreatedAt:     ss.CreatedAt.Format(time.RFC3339),
		},
	})
}
