package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/delivery/http/controller"
	"github.com/rikughi/ecom-api/internal/delivery/http/router"
	"github.com/rikughi/ecom-api/internal/service"

	"github.com/spf13/viper"
)

type App struct {
	Config *viper.Viper
	App    *gin.Engine
}

func Bootstrap(app *App) {
	// setup servie
	helloService := service.NewHelloService()

	// setup controller
	helloController := controller.NewHelloController(helloService)

	router := router.Router{
		App:             app.App,
		HelloController: helloController,
	}

	router.Setup()
}
