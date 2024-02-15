package config

type Config struct {
	BindPort int    `config:"bind-port"`
	BindHost string `config:"bind-host"`
	BaseAddr string `config:"base-addr"`
}
