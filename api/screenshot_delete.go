package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

// ScreenshotDelete is a function to list screenshots
// @Summary Retrieve screenshots
// @Description Retrieve screenshots
// @Param	Token	header	string	true	"Authentication header"
// @Param	limit	query	string	false	"Number of items"
// @Param	page	query	string	false	"Page index"
// @Tags screenshots
// @Produce json
// @Success 200 {object} apimodels.RespScreenshotList{data=[]apimodels.RespScreenshotData}
// @Router /v1/screenshots [get]
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
