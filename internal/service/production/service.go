package production

import (
	"github.com/hell-kitchen/tags/internal/repository"
	"github.com/hell-kitchen/tags/internal/service"
	"go.uber.org/zap"
)

var _ service.TagsService = (*Service)(nil)

type Service struct {
	logger     *zap.Logger
	repository repository.Interface
}

func New(logger *zap.Logger, repo repository.Interface) (*Service, error) {
	srv := &Service{
		logger:     logger,
		repository: repo,
	}
	return srv, nil
}
