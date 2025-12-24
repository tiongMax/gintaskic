package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/tiongMax/gintaskic/internal/config"
	"github.com/tiongMax/gintaskic/internal/router"
)

func main() {
	cfg, err := config.LoadEnvAndGetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":" + cfg.Port)
}
