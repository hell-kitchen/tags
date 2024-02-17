package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/tags/internal/models/dto"
)

type Interface interface {
	Get(ctx context.Context, id uuid.UUID) (*dto.TagDTO, error)
	GetAll(ctx context.Context) ([]dto.TagDTO, error)
	Create(ctx context.Context, tag *dto.TagDTO) error
	CreateMany(ctx context.Context, create []dto.TagDTO) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, tag *dto.TagDTO) error
}
