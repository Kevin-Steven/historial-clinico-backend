package config

import "os"

type Config struct {
	DSN       string
	HTTPPort  string
	JWTSecret string
}

func Load() Config {
	cfg := Config{
		DSN:       os.Getenv("EVOLUCION_DB_DSN"),
		HTTPPort:  os.Getenv("EVOLUCION_HTTP_PORT"),
		JWTSecret: os.Getenv("EVOLUCION_JWT_SECRET"),
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8085"
	}
	return cfg
}
