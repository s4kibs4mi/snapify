package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/s4kibs4mi/snapify/ent"
	"net/http"
)

// ScreenshotDelete is a function to delete a screenshot
// @Summary Delete a specific screenshot
// @Description Delete a specific screenshot
// @Param	Token	header	string	true	"Authentication header"
// @Tags screenshots
// @Produce json
// @Success 204
// @Router /v1/screenshots/{screenshot_id} [delete]
func (h *handlers) ScreenshotDelete(ctx *fiber.Ctx) error {
	screenshotID := ctx.Params("id", "")
	screenshotUUID, err := uuid.Parse(screenshotID)
	if err != nil {
		return h.Serve(ctx, http.StatusBadRequest, map[string]interface{}{"err": err})
	}

	ssDao, tx, err := h.screenshotDao.Tx(nil)
	if err != nil {
		return h.Serve(ctx, http.StatusBadRequest, map[string]interface{}{"err": err})
	}
	defer tx.Rollback()

	ss, err := ssDao.Get(screenshotUUID)
	if err != nil {
		if ent.IsNotFound(err) {
			return h.Serve(ctx, http.StatusNotFound, map[string]interface{}{"err": err})
		}
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	err = ssDao.Delete(screenshotUUID)
	if err != nil {
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	if ss.StoredPath != nil {
		err := h.storageService.Delete(*ss.StoredPath)
		if err != nil {
			return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
		}
	}

	if err := tx.Commit(); err != nil {
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	return h.Serve(ctx, http.StatusNoContent, fiber.Map{})
}
