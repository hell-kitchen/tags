package config

import (
	"context"
	"github.com/hell-kitchen/pkg/confita"
)

type Config struct {
	BindPort int    `config:"bind-port"`
	BindHost string `config:"bind-host"`
	BaseAddr string `config:"base-addr"`
}

func New() (*Config, error) {
	var cfg Config
	if err := confita.Get().Load(context.Background(), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
