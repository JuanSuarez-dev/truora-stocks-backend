package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CockroachDSN string
	APIToken     string
}

func Load() Config {
	// Intentamos cargar .env (local)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró .env, usando variables de entorno")
	}
	return Config{
		CockroachDSN: os.Getenv("COCKROACH_DSN"),
		APIToken:     os.Getenv("API_TOKEN"),
	}
}
