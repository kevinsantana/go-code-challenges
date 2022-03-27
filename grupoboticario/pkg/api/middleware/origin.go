package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func Origin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.SetUserContext(
			context.WithValue(ctx.UserContext(), "origin", ctx.Get("Origin", "Api")),
		)
		return ctx.Next()
	}
}
