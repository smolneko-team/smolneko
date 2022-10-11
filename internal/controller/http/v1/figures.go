package v1

import (
	"errors"
	"strconv"

	"github.com/smolneko-team/smolneko/internal/model"
	"github.com/smolneko-team/smolneko/internal/usecase"
	"github.com/smolneko-team/smolneko/pkg/logger"

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
		h.Get("/", r.figures)
		h.Get("/:id", r.figure)
	}
}

type figuresResponse struct {
	Figures []model.Figure `json:"data"`
}

func (r *figuresRoutes) figures(c *fiber.Ctx) error {
	var count int

	if c.Query("count") == "" {
		count = 20
	} else if value, err := strconv.Atoi(c.Query("count")); err == nil {
		count = value
	} else {
		r.l.Error(err, "http - v1 - count")
		return errorResponse(c, fiber.StatusBadRequest, "Query parameter 'count' is not an integer.")
	}
	if count <= 0 {
		r.l.Error("count is negative or zero", count, "http - v1 - count")
		return errorResponse(c, fiber.StatusBadRequest, "Query parameter 'count' is negative or zero.")
	}

	var offset int
	if c.Query("offset") == "" {
		offset = 0
	} else if value, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = value
	} else {
		r.l.Error(err, "http - v1 - count")
		return errorResponse(c, fiber.StatusBadRequest, "Query parameter 'offset' is not an integer.")
	}

	figures, err := r.f.Figures(c.UserContext(), count, offset)
	if err != nil {
		r.l.Error(err, "http - v1 - figures")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(figuresResponse{figures})
}

type figureResponse struct {
	Figure model.Figure `json:"data"`
}

func (r *figuresRoutes) figure(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := strconv.Atoi(id); err == nil {
		r.l.Error(errors.New("route parameter 'id' is not a string"), "http - v1 - id")
		return errorResponse(c, fiber.StatusBadRequest, "Route parameter 'id' is not a valid id.")
	}

	figure, err := r.f.Figure(c.UserContext(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - figure")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(figureResponse{figure})
}
