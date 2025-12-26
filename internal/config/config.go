package config

import (
	"os"
)

type Config struct {
	Port     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	MongoURI string
	DBName   string
}

func GetConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	dbConfig := DatabaseConfig{
		MongoURI: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}

	if dbConfig.MongoURI == "" {
		dbConfig.MongoURI = "mongodb://localhost:27017"
	}
	if dbConfig.DBName == "" {
		dbConfig.DBName = "taskdb"
	}

	return &Config{
		Port:     port,
		Database: dbConfig,
	}, nil
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return GetConfig()
}
