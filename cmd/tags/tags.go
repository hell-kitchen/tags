package main

import (
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
		),
		fx.Invoke(),
	)
}
