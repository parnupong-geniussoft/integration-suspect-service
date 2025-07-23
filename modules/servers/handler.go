package servers

import (
	"integration-suspect-service/modules/middlewares"
	_suspectControllers "integration-suspect-service/modules/suspect/controllers"
	_suspectRepositories "integration-suspect-service/modules/suspect/repositories"
	_suspectUsecases "integration-suspect-service/modules/suspect/usecases"
	_ktbClients "integration-suspect-service/pkg/clients/ktb"
	_restyClients "integration-suspect-service/pkg/resty_factory"

	_ "integration-suspect-service/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (s *Server) MapHandlers() error {
	// Swagger UI
	s.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	s.App.Use(middlewares.SystemLoggerMiddleware(*s.Log))
	s.App.Use(middlewares.DbLoggerMiddleware(*s.Log))
	s.App.Use(middlewares.RecoverMiddleware())

	s.App.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "ok"})
	})

	// Group a version
	v1 := s.App.Group("/v1")

	// Public routes
	publicGroup := v1.Group("/integration-api/suspect")

	// Initialize Ktb client
	ktbResty := _restyClients.NewRestyClient(s.Cfg, *s.Log, s.Cfg.Suspect.KtbSuspectListHost)
	ktbClient := _ktbClients.NewKtbClient(ktbResty)

	// Suspect Controller
	suspectRepository := _suspectRepositories.NewSuspectRepository(ktbClient)
	suspectUsecase := _suspectUsecases.NewSuspectUsecase(suspectRepository)
	_suspectControllers.NewSuspectController(publicGroup, suspectUsecase, s.Validate)

	// End point not found response
	s.App.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
