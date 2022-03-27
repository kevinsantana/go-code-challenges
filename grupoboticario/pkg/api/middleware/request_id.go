package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rId := ctx.Get("x-request-id")
		if rId == "" {
			rId = uuid.New().String()
		}

		ctx.SetUserContext(
			context.WithValue(ctx.UserContext(), "RequestID", rId),
		)
		return ctx.Next()
	}
}
