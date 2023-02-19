package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *handlers) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authToken, err := h.parseToken(ctx)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"err": err,
			})
		}

		if h.cfg.AuthToken != authToken {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"err": err,
			})
		}

		return ctx.Next()
	}
}

func (h *handlers) parseToken(ctx *fiber.Ctx) (string, error) {
	token, err := h.parseTokenFromHeader(ctx)
	if err != nil {
		return "", err
	}
	if token != "" {
		return token, nil
	}

	token, err = h.parseTokenFromQuery(ctx)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *handlers) parseTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	authHeader := struct {
		Token string `json:"token"`
	}{}
	err := ctx.ReqHeaderParser(&authHeader)
	if err != nil {
		return "", err
	}
	return authHeader.Token, nil
}

func (h handlers) parseTokenFromQuery(ctx *fiber.Ctx) (string, error) {
	authQuery := struct {
		Token string `json:"token"`
	}{}
	err := ctx.QueryParser(&authQuery)
	if err != nil {
		return "", err
	}
	return authQuery.Token, nil
}
