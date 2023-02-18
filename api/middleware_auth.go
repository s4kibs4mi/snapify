package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s4kibs4mi/snapify/ent"
	"net/http"
)

func (h *handlers) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := struct {
			Token string `json:"token"`
		}{}
		err := ctx.ReqHeaderParser(&authHeader)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"err": err,
			})
		}

		_, err = h.tokenDao.Get(authHeader.Token)
		if err != nil {
			if ent.IsNotFound(err) {
				return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"err": err,
				})
			}
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"err": err,
			})
		}

		return ctx.Next()
	}
}
