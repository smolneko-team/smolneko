package usecase

import (
	"context"
	"fmt"

	"github.com/smolneko-team/smolneko/internal/model"
)

type ImagesUseCase struct {
	repo ImagesRepo
}

func NewImages(r ImagesRepo) *ImagesUseCase {
	return &ImagesUseCase{
		repo: r,
	}
}

func (uc ImagesUseCase) Images(ctx context.Context, id, entity string) (model.Image, error) {
	images, err := uc.repo.GetImagesPathById(ctx, id, entity)
	if err != nil {
		return images, fmt.Errorf("ImagesUseCase - Image - uc.repo.GetImagesPathById: %w", err)
	}

	return images, nil
}
