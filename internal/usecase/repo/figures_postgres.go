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
		Select("id, character_id, name, preview_image, description, type, version, size, height, materials, release_date, manufacturer, links, price, created_at, updated_at, is_nsfw, is_draft").
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
		&figure.PreviewImage,
		&figure.Description,
		&figure.Type,
		&figure.Version,
		&figure.Size,
		&figure.Height,
		&figure.Materials,
		&figure.ReleaseDate,
		&figure.Manufacturer,
		&figure.Links,
		&figure.Price,
		&figure.CreatedAt,
		&figure.UpdatedAt,
		&figure.IsNSFW,
		&figure.IsDraft,
	)
	if err != nil {
		return figure, fmt.Errorf("FiguresRepo - GetFigureById - row.Scan: %w", err)
	}

	return figure, nil
}

func (r *FiguresRepo) GetFigures(ctx context.Context, count int, cursor string) ([]model.Figure, string, error) {
	if count > 50 {
		count = 50
	}

	query := r.Builder.
		Select("id, character_id, name, preview_image, description, type, version, size, height, materials, release_date, manufacturer, links, price, created_at, updated_at, is_nsfw, is_draft").
		From("figures").
		OrderBy("created_at DESC, id DESC").
		Limit(uint64(count))

	if cursor != "" {
		created, id, err := decodeCursor(cursor)
		if err != nil {
			return nil, "", fmt.Errorf("FiguresRepo - GetFigures - decodeCursor : %w", err)
		}

		query = query.Where(sq.LtOrEq{
			"created_at": created,
		})

		query = query.Where(sq.Or{
			sq.Lt{
				"created_at": created,
			},
			sq.Lt{
				"id": id,
			},
		})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, "", fmt.Errorf("FiguresRepo - GetFigures - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, "", fmt.Errorf("FiguresRepo - GetFigures - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	figures := make([]model.Figure, 0, _defaultEntityCap)

	for rows.Next() {
		figure := model.Figure{}

		err = rows.Scan(
			&figure.ID,
			&figure.CharacterID,
			&figure.Name,
			&figure.PreviewImage,
			&figure.Description,
			&figure.Type,
			&figure.Version,
			&figure.Size,
			&figure.Height,
			&figure.Materials,
			&figure.ReleaseDate,
			&figure.Manufacturer,
			&figure.Links,
			&figure.Price,
			&figure.CreatedAt,
			&figure.UpdatedAt,
			&figure.IsNSFW,
			&figure.IsDraft,
		)
		if err != nil {
			return nil, "", fmt.Errorf("FiguresRepo - GetFigures - rows.Scan: %w", err)
		}

		figures = append(figures, figure)
	}

	var nextCursor string
	if len(figures) > 0 {
		nextCursor = encodeCursor(figures[len(figures)-1].CreatedAt, figures[len(figures)-1].ID)
	}

	return figures, nextCursor, nil
}
