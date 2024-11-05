package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/service"
	"github.com/rikughi/ecom-api/internal/util"
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

func (c *UserController) Login(ctx *gin.Context) {
	request := new(model.LoginUserRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		util.HandleValidationErrors(ctx, err)
		return
	}

	response, err := c.UserService.Login(ctx, request)
	if err != nil {
		apiErr, ok := err.(*util.ApiError)
		if ok {
			c.Log.Warnf("Failed to login user: %+v", apiErr)
			ctx.JSON(apiErr.Code, model.ApiResponse[*model.UserResponse]{
				Data:   response,
				Errors: apiErr.Error(),
			})
		} else {
			c.Log.Warnf("Failed to login user: %+v", err)
			ctx.JSON(http.StatusInternalServerError, model.ApiResponse[*model.UserResponse]{
				Data:   response,
				Errors: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.UserResponse]{
		Data: response,
	})
}

func (c *UserController) Register(ctx *gin.Context) {
	request := new(model.RegisterUserRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		util.HandleValidationErrors(ctx, err)
		return
	}

	response, err := c.UserService.Create(ctx, request)
	if err != nil {
		apiErr, ok := err.(*util.ApiError)
		if ok {
			c.Log.Warnf("Failed to register user: %+v", apiErr)
			ctx.JSON(apiErr.Code, model.ApiResponse[*model.UserResponse]{
				Data:   response,
				Errors: apiErr.Error(),
			})
		} else {
			c.Log.Warnf("Failed to register user: %+v", err)
			ctx.JSON(http.StatusInternalServerError, model.ApiResponse[*model.UserResponse]{
				Data:   response,
				Errors: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.UserResponse]{
		Data: response,
	})
}
