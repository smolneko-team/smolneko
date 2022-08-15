package v1

import (
	"smolneko/internal/model"
	"smolneko/internal/usecase"
	"smolneko/pkg/logger"
	"strconv"

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
		h.Get("", r.figures)
		h.Get("/:id", r.figure)
	}
}

type figuresResponse struct {
	Figures []model.Figure `json:"figures"`
}

func (r *figuresRoutes) figures(c *fiber.Ctx) error {
	// Optional query parameter
	var count int

	if c.Query("count") == "" {
		count = 20
	} else if value, err := strconv.Atoi(c.Query("count")); err == nil {
		count = value
	} else {
		r.l.Error(err, "http - v1 - count")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	if count <= 0 {
		r.l.Error("count is negative or zero", count, "http - v1 - count")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	figures, err := r.f.Figures(c.UserContext(), count)
	if err != nil {
		r.l.Error(err, "http - v1 - figures")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(figuresResponse{figures})
}

type figureResponse struct {
	Figure model.Figure `json:"figure"`
}

func (r *figuresRoutes) figure(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		r.l.Error(err, "http - v1 - id")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	figure, err := r.f.Figure(c.UserContext(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - figure")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(figureResponse{figure})
}
