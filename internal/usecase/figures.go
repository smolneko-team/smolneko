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
		return nil, fmt.Errorf("FiguresUseCase - Figures - f.repo.GetFigures: %w", err)
	}

	return figures, nil
}

func (uc FiguresUseCase) Figure(ctx context.Context, id int) (model.Figure, error) {
	figure, err := uc.repo.GetFigureById(ctx, id)
	if err != nil {
		return figure, fmt.Errorf("FiguresUseCase - Figure - f.repo.GetFigure: %w", err)
	}

	return figure, nil
}
