package api

import "github.com/gofiber/fiber/v2"

type IHandlers interface {
	AuthMiddleware() fiber.Handler
	ScreenshotCreate(ctx *fiber.Ctx) error
	ScreenshotGet(ctx *fiber.Ctx) error
	ScreenshotView(ctx *fiber.Ctx) error
	ScreenshotDelete(ctx *fiber.Ctx) error
	ScreenshotList(ctx *fiber.Ctx) error
	Serve(ctx *fiber.Ctx, status int, data interface{}) error
}
