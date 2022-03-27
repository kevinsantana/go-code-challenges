package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Healthz(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}

func Readiness(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}
