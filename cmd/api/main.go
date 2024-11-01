package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/config"
)

func main() {
	app := gin.Default()
	viper := config.NewViper()

	config.Bootstrap(&config.App{
		App:    app,
		Config: viper,
	})

	log.Printf("🚀 Server is listening on port %d", viper.GetInt("PORT"))

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
