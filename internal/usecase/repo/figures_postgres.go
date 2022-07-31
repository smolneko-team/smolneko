package repo

import (
    "context"
    "fmt"

    "smolneko/internal/model"
    "smolneko/pkg/postgres"
)

const _defaultEntityCap = 50

type FiguresRepo struct {
    *postgres.Postgres
}

func New(pg *postgres.Postgres) *FiguresRepo {
    return &FiguresRepo{pg}
}

func (r *FiguresRepo) GetFigures(ctx context.Context, count int) ([]model.Figure, error) {
    if count > 50 {
        count = 50
    }

    sql, _, err := r.Builder.
        Select("id, name").
        From("figures").
        Limit(uint64(count)).
        ToSql()

    if err != nil {
        return nil, fmt.Errorf("FiguresRepo - GetFigures - r.Builder: %w", err)
    }

    rows, err := r.Pool.Query(ctx, sql)
    if err != nil {
        return nil, fmt.Errorf("FiguresRepo - GetFigures - r.Pool.Query: %w", err)
    }
    defer rows.Close()

    entities := make([]model.Figure, 0, _defaultEntityCap)

    for rows.Next() {
        e := model.Figure{}

        err = rows.Scan(&e.ID, &e.Name)
        if err != nil {
            return nil, fmt.Errorf("FiguresRepo - GetFigures - rows.Scan: %w", err)
        }

        entities = append(entities, e)
    }

    return entities, nil
}
