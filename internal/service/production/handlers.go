package production

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/tags/internal/models/dto"
)

func (s *Service) Create(ctx context.Context, creationData dto.TagCreationDTO) (*dto.TagDTO, error) {
	tag := &dto.TagDTO{
		ID:    uuid.New(),
		Name:  creationData.Name,
		Color: creationData.Color,
		Slug:  creationData.Slug,
	}

	if err := s.repository.Create(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *Service) Get(ctx context.Context, id string) (*dto.TagDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetAll(ctx context.Context) ([]dto.TagDTO, error) {
	//TODO implement me
	panic("implement me")
}
