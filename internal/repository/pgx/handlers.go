package pgx

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/tags/internal/models/dto"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

const (
	getQuery = `SELECT t.name, t.color, t.slug
FROM tags t
WHERE t.id = $1;`
	getAllQuery = `SELECT t.id, t.name, t.color, t.slug
FROM tags t;`
	createQuery = `INSERT INTO tags (id, name, slug, color)
VALUES ($1, $2, $3, $4);`
	createManyQuery = `INSERT INTO tags(id, name, slug, color)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING;`
	deleteQuery = `DELETE FROM tags WHERE id = $1;`
	updateQuery = `UPDATE tags
SET name  = $2,
    color = $3,
    slug  = $4
WHERE id = $1;`
)

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*dto.TagDTO, error) {
	r.logger.Debug("got Get call", zap.String("id", id.String()))

	res := &dto.TagDTO{
		ID: id,
	}
	err := r.pool.QueryRow(ctx, getQuery, id).Scan(&res.Name, &res.Color, &res.Slug)
	if err != nil {
		r.logger.Error("error just occurred", zap.Error(err))
		return nil, err
	}
	r.logger.Debug("succeed Get call")
	return res, nil
}

func (r *Repository) GetAll(ctx context.Context) (res []dto.TagDTO, err error) {
	var rows pgx.Rows

	rows, err = r.pool.Query(ctx, getAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp dto.TagDTO

		err = rows.Scan(&temp.ID, &temp.Name, &temp.Color, &temp.Slug)
		if err != nil {
			return nil, err
		}

		res = append(res, temp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}

func (r *Repository) Create(ctx context.Context, tag *dto.TagDTO) error {
	_, err := r.pool.Exec(ctx, createQuery, tag.ID, tag.Name, tag.Slug, tag.Color)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateMany(ctx context.Context, tags []dto.TagDTO) error {
	for _, tag := range tags {
		_, err := r.pool.Exec(ctx, createManyQuery, tag.ID, tag.Name, tag.Slug, tag.Color)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, tag *dto.TagDTO) error {
	_, err := r.pool.Exec(ctx, updateQuery, tag.ID, tag.Name, tag.Color, tag.Slug)
	if err != nil {
		return err
	}
	return nil
}
