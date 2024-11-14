package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/delivery/http/controller"
)

type Router struct {
	App               *gin.Engine
	UserController    *controller.UserController
	ProductController *controller.ProductController
	AuthMiddleware    gin.HandlerFunc
}

func (r *Router) Setup() {
	r.App.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Ecom API",
		})
	})

	r.App.POST("/api/users", r.UserController.Register)
	r.App.POST("/api/users/_login", r.UserController.Login)

	r.App.GET("/api/products", r.ProductController.List)
	r.App.GET("/api/products/:id", r.ProductController.Get)

	r.App.Use(r.AuthMiddleware)
	r.App.GET("/api/users/_current", r.UserController.Current)

	r.App.POST("/api/products", r.ProductController.Create)
}
