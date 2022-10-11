package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/smolneko-team/smolneko/internal/model"
	"github.com/smolneko-team/smolneko/pkg/postgres"
)

type CharactersRepo struct {
	*postgres.Postgres
}

func NewCharactersRepo(pg *postgres.Postgres) *CharactersRepo {
	return &CharactersRepo{pg}
}

func (r *CharactersRepo) GetCharacterById(ctx context.Context, id string) (model.Character, error) {
	character := model.Character{}

	sql, _, err := r.Builder.
		Select("id, name, description, birth_at, created_at, updated_at, is_draft").
		From("characters").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return character, fmt.Errorf("CharactersRepo - GetCharacterById - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, id)
	if err != nil {
		return character, fmt.Errorf("CharactersRepo - GetCharacterById - r.Pool.QueryRow: %w", err)
	}

	err = row.Scan(
		&character.ID,
		&character.Name,
		&character.Description,
		&character.BirthAt,
		&character.CreatedAt,
		&character.UpdatedAt,
		&character.IsDraft,
	)
	if err != nil {
		return character, fmt.Errorf("CharactersRepo - GetCharacterById - row.Scan: %w", err)
	}

	return character, nil
}

func (r *CharactersRepo) GetCharacters(ctx context.Context, count, offset int) ([]model.Character, error) {
	if count > 50 {
		count = 50
	}

	query := r.Builder.
		Select("id, name, description, birth_at, created_at, updated_at, is_draft").
		From("characters").
		OrderBy("created_at ASC").
		Limit(uint64(count)).
		Offset(uint64(offset))

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("CharactersRepo - GetCharacters - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("CharactersRepo - GetCharacters - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	characters := make([]model.Character, 0, _defaultEntityCap)

	for rows.Next() {
		character := model.Character{}

		err = rows.Scan(
			&character.ID,
			&character.Name,
			&character.Description,
			&character.BirthAt,
			&character.CreatedAt,
			&character.UpdatedAt,
			&character.IsDraft,
		)
		if err != nil {
			return nil, fmt.Errorf("CharactersRepo - GetCharacters - rows.Scan: %w", err)
		}

		characters = append(characters, character)
	}

	return characters, nil
}
