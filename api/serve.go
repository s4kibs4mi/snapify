package api

import "github.com/gofiber/fiber/v2"

func (h *handlers) Serve(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(data)
}
