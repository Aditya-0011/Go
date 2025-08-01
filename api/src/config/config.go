package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBPath string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("missing required environment variable: DB_PATH")
	}

	return Config{
		Port:   port,
		DBPath: dbPath,
	}
}
