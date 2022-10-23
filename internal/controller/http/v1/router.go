package v1

import (
	prom "github.com/ansrivas/fiberprometheus/v2"
	"github.com/smolneko-team/smolneko/internal/usecase"
	"github.com/smolneko-team/smolneko/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(app *fiber.App, l logger.Interface, f usecase.Figure, c usecase.Character) {
	prometheus := prom.New("smolneko")
	prometheus.RegisterAt(app, "/metrics")

	app.Use(
		fRecover.New(),
		cors.New(cors.ConfigDefault),
		fLogger.New(fLogger.ConfigDefault),
		prometheus.Middleware,
	)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	h := app.Group("/v1")
	{
		newFiguresRoutes(h, f, l)
		newCharactersRoutes(h, c, l)
	}
}
