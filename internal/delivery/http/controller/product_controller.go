package controller

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/service"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Log            *logrus.Logger
	ProductService *service.ProductService
}

func NewProductController(logger *logrus.Logger, productService *service.ProductService) *ProductController {
	return &ProductController{
		Log:            logger,
		ProductService: productService,
	}
}

func (c *ProductController) List(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		util.HandleValidationErrors(ctx, errors.New("Invalid page"))
		return
	}

	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		util.HandleValidationErrors(ctx, errors.New("Invalid Size"))
		return
	}

	query := &model.SearchProductQuery{
		Name: ctx.DefaultQuery("name", ""),
		Page: page,
		Size: size,
	}

	responses, total, err := c.ProductService.Search(ctx, query)
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to search product: %+v", c.Log)
		return
	}

	paging := &model.PageMetadata{
		Page:      query.Page,
		Size:      query.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(query.Size))),
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[[]model.ProductResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *ProductController) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.HandleValidationErrors(ctx, errors.New("Invalid id"))
		return
	}

	response, err := c.ProductService.Get(ctx, int64(id))
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to get product: %+v", c.Log)
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.ProductResponse]{
		Data: response,
	})
}

func (c *ProductController) Create(ctx *gin.Context) {
	request := new(model.CreateProductRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		util.HandleValidationErrors(ctx, err)
		return
	}

	response, err := c.ProductService.Create(ctx, request)
	if err != nil {
		util.HandleApiError(ctx, err, "Failed to create product: %+v", c.Log)
		return
	}

	ctx.JSON(http.StatusOK, model.ApiResponse[*model.ProductResponse]{
		Data: response,
	})
}
