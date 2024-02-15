package production

import (
	"context"
	"github.com/hell-kitchen/tags/internal/models/dto"
)

func (s *Service) Create(ctx context.Context, dto dto.TagCreationDTO) (*dto.TagDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Get(ctx context.Context, id string) (*dto.TagDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetAll(ctx context.Context) ([]dto.TagDTO, error) {
	//TODO implement me
	panic("implement me")
}
