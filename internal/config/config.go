package config

import (
	"os"
	"time"
)

type Config struct {
	DBTimeout time.Duration
}

func Load() *Config {
	timeout := 5 * time.Second // значение по умолчанию
	if v := os.Getenv("DB_TIMEOUT"); v != "" {
		if parsed, err := time.ParseDuration(v); err == nil {
			timeout = parsed
		}
	}
	return &Config{
		DBTimeout: timeout,
	}
}
