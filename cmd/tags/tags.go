package main

import (
	"github.com/hell-kitchen/pkg/logger"
	"github.com/hell-kitchen/tags/internal/config"
	"github.com/hell-kitchen/tags/internal/controller/grpc"
	"github.com/hell-kitchen/tags/internal/service/production"
	"go.uber.org/fx"
)

func main() {
	fx.New(NewOptions()).Run()
}

func NewOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			logger.NewProduction,
			config.NewController,
			grpc.New,
			production.New,
		),
		fx.Invoke(
			startServer,
		),
	)
}

func startServer(lc fx.Lifecycle, ctrl *grpc.Controller) {
	lc.Append(fx.Hook{
		OnStart: ctrl.Start,
		OnStop:  ctrl.Stop,
	})
}
