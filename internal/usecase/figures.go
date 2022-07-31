package usecase

import (
    "context"
    "fmt"

    "smolneko/internal/model"
)

type FiguresUseCase struct {
    repo FiguresRepo
}

func New(r FiguresRepo) *FiguresUseCase {
    return &FiguresUseCase{
        repo: r,
    }
}

func (uc FiguresUseCase) Figures(ctx context.Context, count int) ([]model.Figure, error) {
    figures, err := uc.repo.GetFigures(ctx, count)
    if err != nil {
        return nil, fmt.Errorf("FiguresUseCase - Figures - f.repo.GetFugires: %w", err)
    }

    return figures, nil
}