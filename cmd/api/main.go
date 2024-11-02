package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/config"
)

func main() {
	app := gin.Default()
	viper := config.NewViper()
	log := config.NewLogger(viper)
	db := config.NewDatabase(viper, log)

	config.Bootstrap(&config.App{
		App:    app,
		DB:     db,
		Config: viper,
		Log:    log,
	})

	log.Printf("🚀 Server running on [http://127.0.0.1:%d].", viper.GetInt("PORT"))

	if err := app.Run(fmt.Sprintf(":%d", viper.GetInt("PORT"))); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
