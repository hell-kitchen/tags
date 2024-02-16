package config

import (
	"context"
	"fmt"
	"github.com/hell-kitchen/pkg/confita"
)

type Postgres struct {
	Host     string `config:"postgres-host"`
	Port     int    `config:"postgres-port"`
	User     string `config:"postgres-user"`
	Password string `config:"postgres-password"`
	Database string `config:"postgres-database"`
}

func NewPostgres() (*Postgres, error) {
	cfg := &Postgres{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	}
	if err := confita.Get().Load(context.Background(), cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cfg Postgres) ConnString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)
}
