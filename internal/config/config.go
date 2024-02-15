package config

import (
	"context"
	"fmt"
	"github.com/hell-kitchen/pkg/confita"
)

type Config struct {
	BindPort int    `config:"bind-port,short=p"`
	BindHost string `config:"bind-host,short=h"`
	BaseAddr string `config:"base-addr,short=a"`
	UseTLS   bool   `config:"use-tls"`
	CertFile string `config:"cert-file"`
	KeyFile  string `config:"key-file"`
}

// New initializes new server config.
func New() (*Config, error) {
	cfg := &Config{
		BindPort: 8080,
		BindHost: "0.0.0.0",
		BaseAddr: "http://localhost:8080",
		UseTLS:   false,
		CertFile: "",
		KeyFile:  "",
	}
	if err := confita.Get().Load(context.Background(), cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cfg Config) Bind() string {
	return fmt.Sprintf("%s:%d", cfg.BindHost, cfg.BindPort)
}