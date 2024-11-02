package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/service"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log         *logrus.Logger
	UserService *service.UserService
}

func NewUserController(logger *logrus.Logger, userService *service.UserService) *UserController {
	return &UserController{
		Log:         logger,
		UserService: userService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	request := new(model.RegisterUserRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := make(map[string]string)

			for _, fieldError := range errs {
				validationErrors[fieldError.Field()] = fmt.Sprintf("Field %s: %s", fieldError.Field(), fieldError.Tag())
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": validationErrors,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	response, err := c.UserService.Create(ctx, request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(200, response)
}
