package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Import this
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	// 1. Tell Go to load variables from the .env file into the environment
	// We ignore the error because in production (like Docker/Render), 
	// variables are set manually and there is no .env file.
	_ = godotenv.Load()

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	if cfg.DBHost == "" {
		log.Fatal("Missing DB config: DB_HOST is empty")
	}

	return cfg
}