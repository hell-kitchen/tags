package pgx

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/tags/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

const (
	getQuery = `SELECT t.name, t.color, t.slug
FROM tags t
WHERE t.id = $1;`
	getAllQuery = `SELECT t.id, t.name, t.color, t.slug
FROM tags t;`
	createQuery = `INSERT INTO tags (id, name, slug, color)
VALUES ($1, $2, $3, $4);`
)

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*dto.TagDTO, error) {
	res := &dto.TagDTO{
		ID: id,
	}
	err := r.pool.QueryRow(ctx, getQuery, id).Scan(&res.Name, &res.Color, &res.Slug)
	if err != nil {
		return nil, err
	}
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

func (r *Repository) CreateMany(ctx context.Context, create []dto.TagDTO) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(ctx context.Context, tag *dto.TagDTO) error {
	//TODO implement me
	panic("implement me")
}
