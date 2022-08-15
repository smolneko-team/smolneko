package usecase

import (
	"context"

	"smolneko/internal/model"
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
)
