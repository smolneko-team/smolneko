package v1

import (
	"github.com/smolneko-team/smolneko/internal/usecase"
	"github.com/smolneko-team/smolneko/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(app *fiber.App, l logger.Interface, f usecase.Figure, c usecase.Character) {
	// TODO Config middlewares
	app.Use(
		fRecover.New(),
		cors.New(),
		fLogger.New(),
		// limiter.New(limiter.Config{
		// 	Next: func(c *fiber.Ctx) bool {
		// 		return c.IP() == "127.0.0.1"
		// 	},
		// 	Max: 240,
		// 	KeyGenerator: func(c *fiber.Ctx) string {
		// 		// return c.Get("x-forwarded-for")
		// 		return c.IP()
		// 	},
		// 	LimitReached: func(c *fiber.Ctx) error {
		// 		return c.SendStatus(fiber.StatusTooManyRequests)
		// 	},
		// 	Expiration:        30 * time.Second,
		// 	LimiterMiddleware: limiter.SlidingWindow{},
		// }),
	)

	// app.Get("/dashboard", monitor.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	h := app.Group("/v1")
	{
		newFiguresRoutes(h, f, l)
		newCharactersRoutes(h, c, l)
	}
}
