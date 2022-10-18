package usecase

import (
	"context"
	"fmt"

	"github.com/smolneko-team/smolneko/internal/model"
)

type CharactersUseCase struct {
	repo CharactersRepo
}

func NewCharacters(r CharactersRepo) *CharactersUseCase {
	return &CharactersUseCase{
		repo: r,
	}
}

func (uc CharactersUseCase) Characters(ctx context.Context, count int, cursor string) ([]model.Character, string, error) {
	characters, next, err := uc.repo.GetCharacters(ctx, count, cursor)
	if err != nil {
		return nil, "", fmt.Errorf("CharactersUseCase - Characters - f.repo.GetCharacters: %w", err)
	}

	return characters, next, nil
}

func (uc CharactersUseCase) Character(ctx context.Context, id string) (model.Character, error) {
	character, err := uc.repo.GetCharacterById(ctx, id)
	if err != nil {
		return character, fmt.Errorf("CharactersUseCase - Character - f.repo.GetCharacter: %w", err)
	}

	return character, nil
}
