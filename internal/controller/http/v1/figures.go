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
	f   usecase.Figure
	img usecase.Images
	l   logger.Interface
}

func newFiguresRoutes(handler fiber.Router, f usecase.Figure, img usecase.Images, l logger.Interface) {
	r := &figuresRoutes{f, img, l}

	h := handler.Group("/figures")
	{
		h.Get("/", r.figures)
		h.Get("/:id", r.figure)
		h.Get("/:id/images", r.figureImages)
	}
}

type figuresResponse struct {
	Figures    []model.Figure `json:"data"`
	NextCursor string         `json:"next_cursor"`
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

	cursor := c.Query("cursor")
	figures, next, err := r.f.Figures(c.UserContext(), count, cursor)
	if err != nil {
		r.l.Error(err, "http - v1 - figures")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(figuresResponse{figures, next})
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

func (r *figuresRoutes) figureImages(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := strconv.Atoi(id); err == nil {
		r.l.Error(errors.New("route parameter 'id' is not a string"), "http - v1 - id")
		return errorResponse(c, fiber.StatusBadRequest, "Route parameter 'id' is not a valid id.")
	}

	images, err := r.img.Images(c.UserContext(), id, "figures")
	if err != nil {
		r.l.Error(err, "http - v1 - figure")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(imagesResponse{images})
}
