package v1

import (
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
		h.Get("", r.characters)
		h.Get("/:id", r.character)
	}
}

type charactersResponse struct {
	Characters []model.Character `json:"characters"`
}

func (r *charactersRoutes) characters(c *fiber.Ctx) error {
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

	characters, err := r.c.Characters(c.UserContext(), count)
	if err != nil {
		r.l.Error(err, "http - v1 - figures")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(charactersResponse{characters})
}

type characterResponse struct {
	Character model.Character `json:"character"`
}

func (r *charactersRoutes) character(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		r.l.Error(err, "http - v1 - id")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	character, err := r.c.Character(c.UserContext(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - figure")

		return errorResponse(c, fiber.StatusInternalServerError, "Internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(characterResponse{character})
}
