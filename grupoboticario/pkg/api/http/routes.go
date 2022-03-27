package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/api/handlers"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
	Public      bool
	Scopes      []string
}

var healthCheck = Routes{
	{
		Name:        "Healthz",
		Method:      http.MethodGet,
		Pattern:     "/healthz",
		HandlerFunc: handlers.Healthz,
		Public:      true,
	},
}

func Router() *fiber.App {
	r := fiber.New(fiber.Config{
		Prefork:               false,
		CaseSensitive:         false,
		StrictRouting:         false,
		ServerHeader:          "*",
		AppName:               "Cashback Api",
		Immutable:             true,
		DisableStartupMessage: true,
		ErrorHandler:          middleware.ErrorHandler(),
	})
}
