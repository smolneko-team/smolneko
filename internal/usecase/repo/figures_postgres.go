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
		Select("id, character_id, name, description, type, version, size, height, " +
			"materials, release_date, manufacturer, links, price, created_at, updated_at, is_nsfw, is_draft").
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

func (r *FiguresRepo) GetFigures(ctx context.Context, count int, cursor string) ([]model.Figure, string, string, error) {
	if count > 50 {
		count = 50
	}

	lang := "en"

	columns := r.Builder.Select("id, character_id, name, type, version, size, height, "+
		"materials, release_date, manufacturer, links, price, created_at, updated_at, is_nsfw, is_draft").
		Column("COALESCE(description -> ? #>>'{}', description -> ? #>>'{}') as description", lang, "en").
		From("figures")

	query := r.Builder.Select("figures_cols.*")

	if cursor != "" {
		created, id, suffix, err := decodeCursor(cursor)
		if err != nil {
			return nil, "", "", fmt.Errorf("FiguresRepo - GetFigures - decodeCursor : %w", err)
		}

		if suffix == "next" {
			columns = columns.
				Where(sq.LtOrEq{
					"created_at": created,
				}).Where(sq.Or{
				sq.Lt{
					"created_at": created,
				},
				sq.Lt{
					"id": id,
				},
			}).OrderBy("created_at DESC, id DESC")
		} else {
			columns = columns.
				Where(sq.GtOrEq{
					"created_at": created,
				}).Where(sq.Or{
				sq.Gt{
					"created_at": created,
				},
				sq.Gt{
					"id": id,
				},
			}).OrderBy("created_at ASC, id ASC")
			query = query.OrderBy("created_at DESC, id DESC")
		}
	} else {
		columns = columns.OrderBy("created_at DESC, id DESC")
	}
	columns = columns.Limit(uint64(count))

	query = query.FromSelect(columns, "figures_cols")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, "", "", fmt.Errorf("FiguresRepo - GetFigures - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, "", "", fmt.Errorf("FiguresRepo - GetFigures - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	figures := make([]model.Figure, 0, _defaultEntityCap)

	for rows.Next() {
		figure := model.Figure{}

		err = rows.Scan(
			&figure.ID,
			&figure.CharacterID,
			&figure.Name,
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
			&figure.Description,
		)
		if err != nil {
			return nil, "", "", fmt.Errorf("FiguresRepo - GetFigures - rows.Scan: %w", err)
		}

		figures = append(figures, figure)
	}

	var previousCursor string
	var nextCursor string
	if len(figures) > 0 {
		nextCursor = encodeCursor(figures[len(figures)-1].CreatedAt, figures[len(figures)-1].ID, "next")

		if cursor != "" {
			previousCursor = encodeCursor(figures[0].CreatedAt, figures[0].ID, "prev")
		}
	}

	return figures, nextCursor, previousCursor, nil
}
