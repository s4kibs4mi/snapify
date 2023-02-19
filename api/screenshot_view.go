package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/ent"
	"net/http"
)

// ScreenshotView is a function to view screenshot as PNG
// @Summary Serves screenshot as PNG
// @Description Serves screenshot as PNG
// @Param	Token	query	string	true	"Authentication header"
// @Param	screenshot_id	path	string	true	"Screenshot UUID"
// @Tags screenshots
// @Produce png
// @Success 200
// @Router /v1/screenshots/{screenshot_id}/view [get]
func (h *handlers) ScreenshotView(ctx *fiber.Ctx) error {
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

	object, err := h.storageService.Fetch(*ss.StoredPath)
	if err != nil {
		h.logger.Error(err)
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	return ctx.Status(http.StatusOK).SendStream(object)
}
