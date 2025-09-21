package config

import (
	"log"
	"os"
	"time"
)

type JWTConfig struct {
	Secret     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

type Config struct {
	Environment string
	DatabaseURL string
	DBTimeout   time.Duration
	JWT         JWTConfig
}

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	dbTimeout := getDurationEnv("DB_TIMEOUT", 5*time.Second)
	accessTTL := getDurationEnv("JWT_ACCESS_TTL", 15*time.Minute)
	refreshTTL := getDurationEnv("JWT_REFRESH_TTL", 720*time.Hour)

	return &Config{
		Environment: env,
		DatabaseURL: databaseURL,
		DBTimeout:   dbTimeout,
		JWT: JWTConfig{
			Secret:     jwtSecret,
			AccessTTL:  accessTTL,
			RefreshTTL: refreshTTL,
		},
	}
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
		log.Printf("invalid duration for %s: %v, using default %s", key, err, defaultValue)
	}

	return defaultValue
}
