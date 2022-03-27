package middleware

import (
	"bufio"
	"os"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var API_VERSION string

func init() {
	file, err := os.Open("./.version")
	if err != nil {
		log.WithError(err).Panic("Error to open version file")
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.WithError(err).Panic("Error to close version file")
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		API_VERSION = scanner.Text()
		break
	}

	if API_VERSION == "" {
		log.WithField("api_version", API_VERSION).Panic("Api version is empty")
	}

	log.WithField("api_version", API_VERSION).Info("Api version")
}

func ApiVersion() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		ctx.Set("Api-Version", API_VERSION)
		return err
	}
}
