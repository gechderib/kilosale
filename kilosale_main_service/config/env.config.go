package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		PORT:        getEnv("PORT", "8080"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "password"),
		DB_NAME:     getEnv("DB_NAME", "kilosale"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
