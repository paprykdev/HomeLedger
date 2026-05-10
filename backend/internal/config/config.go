package config

import (
	"os"
)

type Config struct {
	JWTSecret string
	Port      string
	DbPath    string
}

func Load() *Config {
	return &Config{
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Port:      getEnv("PORT", "8080"),
		DbPath:    getEnv("DB_PATH", "./homeledger.db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
