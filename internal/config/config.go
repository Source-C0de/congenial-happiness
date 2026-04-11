package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	//server
	Port        string
	Environment string

	//database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	//jwt
	JWTSecret          string
	JWTExpiration      int
	RefreshExpiryHours int

	//cors
	AllowedOrigins []string
}

func Load() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtEpiry, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY_HOURS"))
	refreshEpiry, _ := strconv.Atoi(os.Getenv("REFRESH_EXPIRY_HOURS"))

	return &Config{
		Port:               getEnv("PORT", "8080"),
		Environment:        getEnv("ENVIRONMENT", "development"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "contacthub_user"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "contacthub_db"),
		DBSSLMode:          getEnv("DB_SSL_MODE", "disable"),
		JWTSecret:          getEnv("JWT_SECRET", "change-this-secret"),
		JWTExpiration:      jwtEpiry,
		RefreshExpiryHours: refreshEpiry,
		AllowedOrigins:     strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:5173"), ","),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
