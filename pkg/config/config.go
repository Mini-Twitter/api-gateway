package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}
	cfg := Config{}

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
