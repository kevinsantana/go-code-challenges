package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/modules/auth"
)

func Auth(scopes []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		key, isCookie := getKey(ctx)
		if key == "" {
			return auth.ErrForbidden
		}
		user, err :=
	}
}