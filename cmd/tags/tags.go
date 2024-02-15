package main

import (
	"github.com/hell-kitchen/tags/internal/config"
	"github.com/hell-kitchen/tags/internal/controller/grpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(NewOptions()).Run()
}

func NewOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			zap.NewProduction,
			config.New,
			grpc.New,
		),
		fx.Invoke(
			startServer,
		),
	)
}

func startServer(lc fx.Lifecycle, ctrl grpc.Controller) {
	lc.Append(fx.Hook{
		OnStart: ctrl.Start,
		OnStop:  ctrl.Stop,
	})
}
