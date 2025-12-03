package config

import (
	"os"
	"sync"
)

type config struct {
	dbPath string
}

var (
	instance *config
	once     sync.Once
)

func Load() *config {
	once.Do(func() {
		instance = &config{
			dbPath: getEnv("DB_PATH", "pkmc.db"),
		}
	})
	return instance
}

func Get() *config {
	if instance == nil {
		return Load()
	}
	return instance
}

func (c *config) GetDBPath() string {
	return c.dbPath
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
