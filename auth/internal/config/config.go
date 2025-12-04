package config

import "os"

type Config struct {
	DSN       string
	HTTPPort  string
	JWTSecret string
}

func Load() Config {
	cfg := Config{
		DSN:       os.Getenv("AUTH_DB_DSN"),
		HTTPPort:  os.Getenv("AUTH_HTTP_PORT"),
		JWTSecret: os.Getenv("AUTH_JWT_SECRET"),
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8082"
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "dev-secret-change-me" // solo desarrollo
	}
	return cfg
}
