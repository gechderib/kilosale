package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT             string
	DB_HOST          string
	DB_INTERNAL_PORT string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_NAME_MAIN     string
	DB_NAME_MONEY    string
	DB_NAME_KYC      string
}

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		PORT:             getEnv("PORT", "8080"),
		DB_HOST:          getEnv("DB_HOST", "localhost"),
		DB_INTERNAL_PORT: getEnv("DB_INTERNAL_PORT", getEnv("DB_PORT", "5432")),
		DB_PORT:          getEnv("DB_PORT", "5432"),
		DB_USER:          getEnv("DB_USER", "postgres"),
		DB_PASSWORD:      getEnv("DB_PASSWORD", "password"),
		DB_NAME_MAIN:     getEnv("DB_NAME_MAIN", "kilosale_db"),
		DB_NAME_MONEY:    getEnv("DB_NAME_MONEY", "kilosale_money_db"),
		DB_NAME_KYC:      getEnv("DB_NAME_KYC", "kilosale_kyc_db"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
