package v1

import (
	"github.com/smolneko-team/smolneko/internal/usecase"
	"github.com/smolneko-team/smolneko/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(app *fiber.App, l logger.Interface, f usecase.Figure, c usecase.Character, img usecase.Images) {
	app.Use(
		recover.New(),
		cors.New(cors.ConfigDefault),
		fLogger.New(fLogger.ConfigDefault),
	)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	h := app.Group("/v1")
	{
		newFiguresRoutes(h, f, img, l)
		newCharactersRoutes(h, c, img, l)
	}
}
