package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/api/handlers"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/api/middleware"
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

	r.Use(
		middleware.RequestID(),
		middleware.CORS(),
		middleware.Recover(),
		compress.New(compress.Config{
			Level: compress.LevelDefault,
		}),
		middleware.Origin(),
		middleware.ApiVersion(),
	)

	api := r.Group("/api")
	for _, route := range healthCheck {
		api.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	v1 := api.Group("/v1")

	var routes []Route
	routes = append(routes, healthCheck...)

	for _, route := range routes {
		if route.Public {
			v1.Add(route.Method, route.Pattern, route.HandlerFunc)
			continue
		}
		v1.Add(route.Method, route.Pattern, middleware.Auth(route.Scopes), route.HandlerFunc)
	}

	r.Use(middleware.RouteNotFound())

	return r
}
