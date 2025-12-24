package database

// import (
// 	"fmt"
// 	"log"

// 	"github.com/tiongMax/gintaskic/internal/config"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect(cfg config.DatabaseConfig) error {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
// 		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to database: %w", err)
// 	}

// 	DB = db
// 	log.Println("Successfully connected to the database")
// 	return nil
// }
