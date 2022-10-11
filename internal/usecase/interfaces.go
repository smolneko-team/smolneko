package usecase

import (
	"context"

	"github.com/smolneko-team/smolneko/internal/model"
)

type (
	Figure interface {
		Figures(context.Context, int, int) ([]model.Figure, error)
		Figure(context.Context, string) (model.Figure, error)
	}

	FiguresRepo interface {
		GetFigures(context.Context, int, int) ([]model.Figure, error)
		GetFigureById(context.Context, string) (model.Figure, error)
	}

	Character interface {
		Characters(context.Context, int, int) ([]model.Character, error)
		Character(context.Context, string) (model.Character, error)
	}

	CharactersRepo interface {
		GetCharacters(context.Context, int, int) ([]model.Character, error)
		GetCharacterById(context.Context, string) (model.Character, error)
	}
)
