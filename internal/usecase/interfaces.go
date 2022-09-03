package usecase

import (
	"context"

	"github.com/smolneko-team/smolneko/internal/model"
)

type (
	Figure interface {
		Figures(context.Context, int) ([]model.Figure, error)
		Figure(context.Context, int) (model.Figure, error)
	}

	FiguresRepo interface {
		GetFigures(context.Context, int) ([]model.Figure, error)
		GetFigureById(context.Context, int) (model.Figure, error)
	}

	Character interface {
		Characters(context.Context, int) ([]model.Character, error)
		Character(context.Context, int) (model.Character, error)
	}

	CharactersRepo interface {
		GetCharacters(context.Context, int) ([]model.Character, error)
		GetCharacterById(context.Context, int) (model.Character, error)
	}
)
