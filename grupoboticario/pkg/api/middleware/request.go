package middleware

import "github.com/gofiber/fiber/v2"

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code, res, isLog := handlers.GetResponseError(err)
	}
}
