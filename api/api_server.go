package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/s4kibs4mi/snapify/config"
)

type Server struct {
	FiberApp  *fiber.App
	AppConfig *config.AppCfg
}

func NewServer(cfg *config.AppCfg) *Server {
	server := &Server{
		AppConfig: cfg,
	}

	server.FiberApp = fiber.New()
	return server
}

func (s *Server) App() *fiber.App {
	return s.FiberApp
}

func (s *Server) Start(handlers IHandlers) error {
	s.FiberApp.Use(recover2.New())
	s.FiberApp.Use(logger.New())
	s.FiberApp.Get("/docs/*", swagger.HandlerDefault)

	v1 := s.FiberApp.Group("/v1")
	v1.Use(handlers.AuthMiddleware())
	v1.Post("/screenshots/", handlers.ScreenshotCreate)
	v1.Get("/screenshots/:id/", handlers.ScreenshotGet)
	v1.Get("/screenshots/:id/view/", handlers.ScreenshotView)
	v1.Delete("/screenshots/:id/", handlers.ScreenshotDelete)
	v1.Get("/screenshots/", handlers.ScreenshotList)

	return s.FiberApp.Listen(fmt.Sprintf("%s:%d", s.AppConfig.Base, s.AppConfig.Port))
}

func (s *Server) Stop() error {
	return s.FiberApp.Shutdown()
}
