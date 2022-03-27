package middleware

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/api/handlers"
)

func Recover() fiber.Handler {
	return func(ctx *fiber.Ctx) (errHandler error) {
		defer func() {
			if err := recover(); err != nil {
				errHandler = err.(error)
				ctx.Response().SetStatusCode(500)
				request, _ := httputil.DumpRequest(toHTTPRequest(ctx.Context()), false)
				log.
					WithField("stack", string(debug.Stack())).
					WithField("request", string(request)).
					WithError(err.(error)).
					Error("Panic recovered")
			}
		}()

		if errHandler != nil {
			return errHandler
		}

		return ctx.Next()
	}
}

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code, res, isLog := handlers.GetResponseError(err)

		if isLog {
			log.
				WithFields(log.Fields{
					"path":          string(ctx.Request().URI().Path()),
					"status_code":   code,
					"client_ip":     ctx.IP(),
					"method":        string(ctx.Context().Method()),
					"user_agent":    string(ctx.Request().Header.UserAgent()),
					"response_body": res,
					"request_body":  string(ctx.Request().Body()),
				}).
				WithError(err).
				Error("Error not map")
		}
		return ctx.Status(code).JSON(res)
	}
}

func RouteNotFound() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).
			JSON(handlers.ResponseError{
				Code:    "API|ROUTE_NOT_FOUND",
				Message: "Route not found",
			})
	}
}

func toHTTPRequest(ctx *fasthttp.RequestCtx) *http.Request {
	uri := ctx.Request.URI()
	requestUrl := &url.URL{
		Scheme:   string(uri.Scheme()),
		Path:     string(uri.Path()),
		Host:     string(uri.Host()),
		RawQuery: string(uri.QueryString()),
	}

	return &http.Request{
		Method: string(ctx.Request.Header.Method()),
		URL:    requestUrl,
		Proto:  "HTTP/1.1",
		Header: transformRequestHeaders(&ctx.Request),
		Host:   string(uri.Host()),
		TLS:    ctx.TLSConnectionState(),
	}
}

func transformRequestHeaders(r *fasthttp.Request) http.Header {
	header := make(http.Header)
	r.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		header.Set(sk, sv)
	})

	return header
}
