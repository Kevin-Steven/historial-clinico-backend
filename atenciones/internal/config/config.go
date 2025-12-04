package config

import "os"

type Config struct {
	DSN       string
	HTTPPort  string
	JWTSecret string
}

func Load() Config {
	cfg := Config{
		DSN:       os.Getenv("ATENCIONES_DB_DSN"),
		HTTPPort:  os.Getenv("ATENCIONES_HTTP_PORT"),
		JWTSecret: os.Getenv("ATENCIONES_JWT_SECRET"),
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8083"
	}
	return cfg
}
