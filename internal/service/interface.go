package service

import (
	"context"
	"github.com/hell-kitchen/tags/internal/models/dto"
	"go.uber.org/zap"
)

type TagsService interface {
	Create(ctx context.Context, dto dto.TagCreationDTO, logFields ...zap.Field) (*dto.TagDTO, error)
	CreateMany(ctx context.Context, create []dto.TagCreationDTO, logFields ...zap.Field) ([]dto.TagDTO, error)
	Get(ctx context.Context, id string, logFields ...zap.Field) (*dto.TagDTO, error)
	GetAll(ctx context.Context, logFields ...zap.Field) ([]dto.TagDTO, error)
	Delete(ctx context.Context, id string, logFields ...zap.Field) error
	Update(ctx context.Context, dto dto.TagUpdateDTO, logFields ...zap.Field) (*dto.TagDTO, error)
}
