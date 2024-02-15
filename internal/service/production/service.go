package production

import (
	"github.com/hell-kitchen/tags/internal/service"
	"go.uber.org/zap"
)

var _ service.TagsService = (*Service)(nil)

type Service struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) (*Service, error) {
	srv := &Service{
		logger: logger,
	}
	return srv, nil
}
