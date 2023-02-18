package api

import (
	"github.com/gofiber/fiber/v2"
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

		if h.cfg.AuthToken != authHeader.Token {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"err": err,
			})
		}

		return ctx.Next()
	}
}
