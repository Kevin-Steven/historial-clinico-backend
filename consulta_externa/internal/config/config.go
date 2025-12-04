package config

import "os"

type Config struct {
	DSN      string
	HTTPPort string
}

func Load() Config {
	cfg := Config{
		DSN:      os.Getenv("CONSULTA_EXT_DB_DSN"),
		HTTPPort: os.Getenv("CONSULTA_EXT_HTTP_PORT"),
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8084"
	}
	return cfg
}
