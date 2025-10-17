package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db := os.Getenv("DATABASE_URL")
	return Config{Port: port, DatabaseURL: db}
}
