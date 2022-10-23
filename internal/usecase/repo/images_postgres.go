package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/smolneko-team/smolneko/internal/model"
	"github.com/smolneko-team/smolneko/pkg/postgres"
)

type ImagesRepo struct {
	*postgres.Postgres
}

func NewImagesRepo(pg *postgres.Postgres) *ImagesRepo {
	return &ImagesRepo{pg}
}

func (r *ImagesRepo) GetImagesPathById(ctx context.Context, id string, entity string) (model.Image, error) {
	images := model.Image{}

	query := r.Builder.Select("id, path")

	if entity == "figures" {
		query = query.
			From("figures_images").
			Where(sq.Eq{"figure_id": id})
	}
	if entity == "characters" {
		query = query.
			From("characters_images").
			Where(sq.Eq{"character_id": id})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return images, fmt.Errorf("ImagesRepo - GetImagesPathById - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(
		&images.ID,
		&images.URL,
	)
	if err != nil {
		return images, fmt.Errorf("ImagesRepo - GetImagesPathById - row.Scan: %w", err)
	}
	images.Count = len(images.URL)

	return images, nil
}
