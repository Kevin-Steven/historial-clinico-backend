package config

import "os"

type Config struct {
	DSN       string
	HTTPPort  string
	JWTSecret string
}

func Load() Config {
	cfg := Config{
		DSN:       os.Getenv("PACIENTES_DB_DSN"),
		HTTPPort:  os.Getenv("PACIENTES_HTTP_PORT"),
		JWTSecret: os.Getenv("PACIENTES_JWT_SECRET"),
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8081"
	}
	return cfg
}
