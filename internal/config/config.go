package config

import (
	"os"
)

type Config struct {
	Port string
}

func GetConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	return &Config{
		Port: port,
	}, nil
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return GetConfig()
}
