package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	MySQLURL      = AddEnvVariable("MYSQL_URL", "")
	RedisUsername = AddEnvVariable("REDIS_USERNAME", "admin")
	RedisPassword = AddEnvVariable("REDIS_PASSWORD", "admin")
)

func AddEnvVariable(variable, fallback string) string {
	if val := os.Getenv(variable); val != "" {
		return val
	}
	return fallback
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("cannot load .env: %w", err))
	}
}
