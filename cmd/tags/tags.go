package main

import (
	"github.com/google/uuid"
	"github.com/hell-kitchen/pkg/logger"
	"github.com/hell-kitchen/pkg/postgres"
	"github.com/hell-kitchen/tags/internal/config"
	"github.com/hell-kitchen/tags/internal/controller/grpc"
	"github.com/hell-kitchen/tags/internal/repository"
	"github.com/hell-kitchen/tags/internal/repository/pgx"
	"github.com/hell-kitchen/tags/internal/service"
	"github.com/hell-kitchen/tags/internal/service/production"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(NewOptions()).Run()
}

func NewOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			getLogger,
			config.NewController,
			config.NewPostgres,
			grpc.New,
			fx.Annotate(production.New, fx.As(new(service.TagsService))),
			NewPool,
			fx.Annotate(pgx.New, fx.As(new(repository.Interface))),
		),
		fx.Invoke(
			startServer,
		),
		fx.NopLogger,
	)
}

func getLogger() (*zap.Logger, error) {
	log, err := logger.NewProduction(
		zap.Fields(
			logger.WithInstanceID(uuid.NewString()),
			logger.WithService("tags"),
		),
	)
	return log, err
}

func startServer(lc fx.Lifecycle, ctrl *grpc.Controller) {
	lc.Append(fx.Hook{
		OnStart: ctrl.Start,
		OnStop:  ctrl.Stop,
	})
}

func NewPool(lc fx.Lifecycle, log *zap.Logger, cfg *config.Postgres) (*pgxpool.Pool, error) {
	return postgres.NewWithFx(lc, cfg.ConnString(), log)
}
