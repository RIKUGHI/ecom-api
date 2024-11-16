package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/service"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
)

type CartController struct {
	Log          *logrus.Logger
	OrderService *service.OrderService
}

func NewCartController(logger *logrus.Logger, OrderService *service.OrderService) *CartController {
	return &CartController{
		Log:          logger,
		OrderService: OrderService,
	}
}

func (c *CartController) Checkout(ctx *gin.Context) {
	request := new(model.CartCheckoutRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		util.HandleValidationErrors(ctx, err)
		return
	}

	if len(request.Items) == 0 {
		util.HandleValidationErrors(ctx, errors.New("Minimum 1 item"))
		return
	}

	response, err := c.OrderService.Create(ctx, request)
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to create order: %+v", c.Log)
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.CheckoutResponse]{
		Data: response,
	})
}
