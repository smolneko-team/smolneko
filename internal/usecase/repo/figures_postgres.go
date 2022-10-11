package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/smolneko-team/smolneko/internal/model"
	"github.com/smolneko-team/smolneko/pkg/postgres"
)

type FiguresRepo struct {
	*postgres.Postgres
}

func NewFiguresRepo(pg *postgres.Postgres) *FiguresRepo {
	return &FiguresRepo{pg}
}

func (r *FiguresRepo) GetFigureById(ctx context.Context, id string) (model.Figure, error) {
	figure := model.Figure{}

	sql, args, err := r.Builder.
		Select("id, character_id, name, description, type, size, height, materials, release_date, manufacturer, links, price, created_at, updated_at, is_draft").
		From("figures").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return figure, fmt.Errorf("FiguresRepo - GetFigureById - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(
		&figure.ID,
		&figure.CharacterID,
		&figure.Name,
		&figure.Description,
		&figure.Type,
		&figure.Size,
		&figure.Height,
		&figure.Materials,
		&figure.ReleaseDate,
		&figure.Manufacturer,
		&figure.Links,
		&figure.Price,
		&figure.CreatedAt,
		&figure.UpdatedAt,
		&figure.IsDraft,
	)
	if err != nil {
		return figure, fmt.Errorf("FiguresRepo - GetFigureById - row.Scan: %w", err)
	}

	return figure, nil
}

func (r *FiguresRepo) GetFigures(ctx context.Context, count int, offset int) ([]model.Figure, error) {
	if count > 50 {
		count = 50
	}

	query := r.Builder.
		Select("id, character_id, name, description, type, size, height, materials, release_date, manufacturer, links, price, created_at, updated_at, is_draft").
		From("figures").
		OrderBy("created_at ASC").
		Limit(uint64(count)).
		Offset(uint64(offset))

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("FiguresRepo - GetFigures - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("FiguresRepo - GetFigures - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	figures := make([]model.Figure, 0, _defaultEntityCap)

	for rows.Next() {
		figure := model.Figure{}

		err = rows.Scan(
			&figure.ID,
			&figure.CharacterID,
			&figure.Name,
			&figure.Description,
			&figure.Type,
			&figure.Size,
			&figure.Height,
			&figure.Materials,
			&figure.ReleaseDate,
			&figure.Manufacturer,
			&figure.Links,
			&figure.Price,
			&figure.CreatedAt,
			&figure.UpdatedAt,
			&figure.IsDraft,
		)
		if err != nil {
			return nil, fmt.Errorf("FiguresRepo - GetFigures - rows.Scan: %w", err)
		}

		figures = append(figures, figure)
	}

	return figures, nil
}
