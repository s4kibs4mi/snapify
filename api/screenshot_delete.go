package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

func (h *handlers) ScreenshotDelete(ctx *fiber.Ctx) error {
	screenshotID := ctx.Params("id", "")
	screenshotUUID, err := uuid.Parse(screenshotID)
	if err != nil {
		return h.Serve(ctx, http.StatusBadRequest, map[string]interface{}{"err": err})
	}

	err = h.screenshotDao.Delete(screenshotUUID)
	if err != nil {
		return h.Serve(ctx, http.StatusInternalServerError, map[string]interface{}{"err": err})
	}

	return h.Serve(ctx, http.StatusNoContent, fiber.Map{})
}
