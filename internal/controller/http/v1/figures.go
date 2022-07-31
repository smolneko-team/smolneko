package v1

import (
    "smolneko/internal/model"
    "smolneko/internal/usecase"
    "smolneko/pkg/logger"

    "github.com/gofiber/fiber/v2"
)

type figuresRoutes struct {
    f usecase.Figure
    l logger.Interface
}

func newFiguresRoutes(handler fiber.Router, f usecase.Figure, l logger.Interface) {
    r := &figuresRoutes{f, l}

    h := handler.Group("/figures")
    {
        h.Get("/:count?", r.figures) // optional
    }
}

type figuresResponse struct {
    Figures []model.Figure `json:"figures"`
}

func (r *figuresRoutes) figures(c *fiber.Ctx) error {
    count, err := c.ParamsInt("count")
    if err != nil {
        r.l.Error(err, "http - v1 - count")

        return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
    }

    figures, err := r.f.Figures(c.UserContext(), count)
    if err != nil {
        r.l.Error(err, "http - v1 - figures")

        return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
    }

    return c.Status(fiber.StatusOK).JSON(figuresResponse{figures})
}
