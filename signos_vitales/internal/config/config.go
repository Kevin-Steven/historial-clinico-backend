package config

import "os"

type Config struct {
	DBDSN     string
	HTTPPort  string
	JWTSecret string
}

func Load() Config {
	return Config{
		DBDSN:     os.Getenv("SIGNOS_DB_DSN"),
		HTTPPort:  os.Getenv("SIGNOS_HTTP_PORT"),
		JWTSecret: os.Getenv("SIGNOS_JWT_SECRET"),
	}
}
