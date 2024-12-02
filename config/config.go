package config

import (
	"os"

	"github.com/joho/godotenv"

	"boilerplate/pkg/logger"
)

type Config struct {
	Port     string
	LogLevel string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

func Load() *Config {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log := logger.Init("error")
		log.Error("Warning: no .env file found, using default values")
	}

	return &Config{
		Port:     getEnv("PORT", ""),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		DBHost:   getEnv("DB_HOST", ""),
		DBPort:   getEnv("DB_PORT", ""),
		DBUser:   getEnv("DB_USER", ""),
		DBPass:   getEnv("DB_PASS", ""),
		DBName:   getEnv("DB_NAME", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
