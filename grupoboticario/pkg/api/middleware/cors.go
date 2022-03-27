package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     strings.Join([]string{cors.ConfigDefault.AllowMethods, "OPTIONS"}, ","),
		AllowHeaders:     "*",
	})
}
