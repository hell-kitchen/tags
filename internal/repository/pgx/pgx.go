package pgx

import (
	"github.com/hell-kitchen/tags/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var _ repository.Interface = (*Repository)(nil)

type Repository struct {
	logger *zap.Logger
	pool   *pgxpool.Pool
}

func New(log *zap.Logger, pool *pgxpool.Pool) (*Repository, error) {
	repo := &Repository{
		logger: log,
		pool:   pool,
	}
	return repo, nil
}
