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
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	result, err := s.repository.Get(ctx, parsedUUID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetAll(ctx context.Context) ([]dto.TagDTO, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) CreateMany(ctx context.Context, create []dto.TagCreationDTO) ([]dto.TagDTO, error) {
	var result = make([]dto.TagDTO, 0, len(create))
	for _, tag := range create {
		temp := dto.TagDTO{
			ID:    uuid.New(),
			Name:  tag.Name,
			Color: tag.Color,
			Slug:  tag.Slug,
		}
		result = append(result, temp)
	}

	err := s.repository.CreateMany(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Delete(ctx context.Context, rawID string) error {
	id, err := uuid.Parse(rawID)
	if err != nil {
		return err
	}

	err = s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, dto dto.TagUpdateDTO) (*dto.TagDTO, error) {
	tag, err := s.Get(ctx, dto.ID)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil && *dto.Name != tag.Name {
		tag.Name = *dto.Name
	}
	if dto.Slug != nil && *dto.Slug != tag.Slug {
		tag.Slug = *dto.Slug
	}
	if dto.Color != nil && *dto.Color != tag.Color {
		tag.Color = *dto.Color
	}

	err = s.repository.Update(ctx, tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
