package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s4kibs4mi/snapify/apimodels"
	"net/http"
	"time"
)

func (h *handlers) ScreenshotCreate(ctx *fiber.Ctx) error {
	req := &apimodels.ReqScreenshotCreate{}
	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusBadRequest, map[string]interface{}{"err": err})
	}
	if err := req.Validate(); err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusUnprocessableEntity, err)
	}

	screenshotDao, tx, err := h.screenshotDao.Tx(nil)
	if err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}
	// Rollback in case of failure
	defer tx.Rollback()

	// Create screenshot
	screenshot, err := screenshotDao.Create(req)
	if err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	// Queue to take screenshot of URL
	err = h.queueService.EnqueueTakeScreenshot(screenshot.ID)
	if err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	if err := tx.Commit(); err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	return h.Serve(ctx, http.StatusAccepted, map[string]interface{}{
		"data": apimodels.RespScreenshot{
			ID:        screenshot.ID.String(),
			URL:       screenshot.URL,
			Status:    string(screenshot.Status),
			CreatedAt: screenshot.CreatedAt.Format(time.RFC3339),
		},
	})
}
