package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/tiongMax/gintaskic/internal/config"
	// "github.com/tiongMax/gintaskic/internal/database"
	"github.com/tiongMax/gintaskic/internal/router"
)

func main() {
	cfg, err := config.LoadEnvAndGetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// if err := database.Connect(cfg.Database); err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	r := gin.Default()
	router.SetupRouter(r, cfg)
	r.Run(":" + cfg.Port)
}
