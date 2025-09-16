package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port      int
	Debug     bool
	JWTSecret string
	// Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSchema   string
}

var App *Config

func Load() {
	_ = godotenv.Load()
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Invalid PORT value: %v", err)
	}
	debug := getEnv("DEBUG", "false") == "true"
	dbPort, err := strconv.Atoi(getEnv("BLUEPRINT_DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT value: %v", err)
	}
	App = &Config{
		Port:      port,
		Debug:     debug,
		JWTSecret: getEnvOrFail("JWT_SECRET"),

		// Database
		DBHost:     getEnvOrFail("BLUEPRINT_DB_HOST"),
		DBPort:     dbPort,
		DBUser:     getEnvOrFail("BLUEPRINT_DB_USERNAME"),
		DBPassword: getEnvOrFail("BLUEPRINT_DB_PASSWORD"),
		DBName:     getEnvOrFail("BLUEPRINT_DB_DATABASE"),
		DBSchema:   getEnvOrFail("BLUEPRINT_DB_SCHEMA"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
func getEnvOrFail(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
