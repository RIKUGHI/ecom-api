package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/delivery/http/controller"
	"github.com/rikughi/ecom-api/internal/delivery/http/middleware"
	"github.com/rikughi/ecom-api/internal/delivery/http/router"
	"github.com/rikughi/ecom-api/internal/repository"
	"github.com/rikughi/ecom-api/internal/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

type App struct {
	App    *gin.Engine
	DB     *gorm.DB
	Config *viper.Viper
	Log    *logrus.Logger
}

func Bootstrap(app *App) {
	// setup repository
	userRepository := repository.NewUserRepository()
	productRepository := repository.NewProductRepository()

	// setup servie
	userService := service.NewUserService(app.DB, app.Log, userRepository)
	productService := service.NewProductService(app.DB, app.Log, productRepository)

	// setup controller
	userController := controller.NewUserController(app.Log, userService)
	productController := controller.NewProductController(app.Log, productService)

	// setup middleware
	authMiddleware := middleware.NewAuth(app.Log)

	router := router.Router{
		App:               app.App,
		UserController:    userController,
		ProductController: productController,
		AuthMiddleware:    authMiddleware,
	}

	router.Setup()
}
