package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/service"
)

type HelloController struct {
	Service *service.HelloService
}

func NewHelloController(service *service.HelloService) *HelloController {
	return &HelloController{
		Service: service,
	}
}

func (h *HelloController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": h.Service.Hello(),
	})
}
