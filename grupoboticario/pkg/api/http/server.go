package http

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app"
)

func Run(cfg app.Config) {
	r := Router()
	ListenAndServe(cfg, r)
}

func ListenAndServe(cfg app.Config, srv *fiber.App) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done
		// log.Info("Gracefully shutting down...")
		_ = srv.Shutdown()
	}()

	// log.Info("Server started")
	srvHost := net.JoinHostPort(cfg.Host, cfg.Port)
	// log.Info("running on %s", srvHost)
	if err := srv.Listen(srvHost); err != nil && err != http.ErrServerClosed {
		log.Panicf("server error: %v", err)
	}
}
