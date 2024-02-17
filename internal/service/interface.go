package service

import (
	"context"
	"github.com/hell-kitchen/tags/internal/models/dto"
)

type TagsService interface {
	Create(ctx context.Context, dto dto.TagCreationDTO) (*dto.TagDTO, error)
	CreateMany(ctx context.Context, create []dto.TagCreationDTO) ([]dto.TagDTO, error)
	Get(ctx context.Context, id string) (*dto.TagDTO, error)
	GetAll(ctx context.Context) ([]dto.TagDTO, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, dto dto.TagUpdateDTO) (*dto.TagDTO, error)
}
