package usecase

import (
    "context"

    "smolneko/internal/model"
)

type (
    Figure interface {
        Figures(context.Context, int) ([]model.Figure, error)
    }

    FiguresRepo interface {
        GetFigures(context.Context, int) ([]model.Figure, error)
    }
)