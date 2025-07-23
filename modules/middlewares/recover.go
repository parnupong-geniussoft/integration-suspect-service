package middlewares

import (
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

func RecoverMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v\n", r)
				debug.PrintStack()
				ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()
		return ctx.Next()
	}
}
