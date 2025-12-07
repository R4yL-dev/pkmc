package config

import (
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	dbPath         string
	defaultTimeout time.Duration
}

var (
	instance *Config
	once     sync.Once
)

func Load() *Config {
	once.Do(func() {
		instance = &Config{
			dbPath:         getEnv("DB_PATH", "pkmc.db"),
			defaultTimeout: getDurationEnv("DEFAULT_TIMEOUT", 30*time.Second),
		}
	})
	return instance
}

func Get() *Config {
	if instance == nil {
		return Load()
	}
	return instance
}

func (c *Config) GetDBPath() string {
	return c.dbPath
}

func (c *Config) GetDefaultTimeout() time.Duration {
	return c.defaultTimeout
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}
