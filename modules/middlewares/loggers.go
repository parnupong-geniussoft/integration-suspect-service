package middlewares

import (
	"integration-suspect-service/pkg/loggers"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SystemLoggerMiddleware(logger loggers.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		logger.SystemLogger(ctx, start, err)
		return err
	}
}

func DbLoggerMiddleware(logger loggers.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := logger.DbLogger(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
