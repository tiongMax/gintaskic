package config

import (
	"os"
)

type Config struct {
	Port string
	// Database DatabaseConfig
}

// type DatabaseConfig struct {
// 	Host     string
// 	User     string
// 	Password string
// 	Name     string
// 	Port     string
// 	SSLMode  string
// }

func GetConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// dbConfig := DatabaseConfig{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Name:     os.Getenv("DB_NAME"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// }

	return &Config{
		Port: port,
		// Database: dbConfig,
	}, nil
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return GetConfig()
}
