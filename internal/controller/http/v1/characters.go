package v1

import (
	"errors"
	"strconv"

	"github.com/smolneko-team/smolneko/internal/model"
	"github.com/smolneko-team/smolneko/internal/usecase"
	"github.com/smolneko-team/smolneko/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type charactersRoutes struct {
	c usecase.Character
	l logger.Interface
}

func newCharactersRoutes(handler fiber.Router, c usecase.Character, l logger.Interface) {
	r := &charactersRoutes{c, l}

	h := handler.Group("/characters")
	{
		h.Get("/", r.characters)
		h.Get("/:id", r.character)
	}
}

type charactersResponse struct {
	Characters []model.Character `json:"data"`
	NextCursor string            `json:"next_cursor"`
}

func (r *charactersRoutes) characters(c *fiber.Ctx) error {
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
	characters, next, err := r.c.Characters(c.UserContext(), count, cursor)
	if err != nil {
		r.l.Error(err, "http - v1 - figures")
		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(charactersResponse{characters, next})
}

type characterResponse struct {
	Character model.Character `json:"data"`
}

func (r *charactersRoutes) character(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := strconv.Atoi(id); err == nil {
		r.l.Error(errors.New("route parameter 'id' is not a string"), "http - v1 - id")
		return errorResponse(c, fiber.StatusBadRequest, "Route parameter 'id' is not a valid id.")
	}

	character, err := r.c.Character(c.UserContext(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - figure")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(characterResponse{character})
}
