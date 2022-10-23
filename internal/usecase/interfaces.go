package usecase

import (
	"context"

	"github.com/smolneko-team/smolneko/internal/model"
)

type (
	Figure interface {
		Figures(context.Context, int, string) ([]model.Figure, string, error)
		Figure(context.Context, string) (model.Figure, error)
	}

	FiguresRepo interface {
		GetFigures(context.Context, int, string) ([]model.Figure, string, error)
		GetFigureById(context.Context, string) (model.Figure, error)
	}

	Character interface {
		Characters(context.Context, int, string) ([]model.Character, string, error)
		Character(context.Context, string) (model.Character, error)
	}

	CharactersRepo interface {
		GetCharacters(context.Context, int, string) ([]model.Character, string, error)
		GetCharacterById(context.Context, string) (model.Character, error)
	}

	Images interface {
		Images(context.Context, string, string) (model.Image, error)
	}

	ImagesRepo interface {
		GetImagesPathById(context.Context, string, string) (model.Image, error)
	}
)
