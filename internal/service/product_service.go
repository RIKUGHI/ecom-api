package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/model/converter"
	"github.com/rikughi/ecom-api/internal/repository"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	ProductRepository *repository.ProductRepository
}

func NewProductService(DB *gorm.DB, Log *logrus.Logger, ProductRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		DB:                DB,
		Log:               Log,
		ProductRepository: ProductRepository,
	}
}

func (s *ProductService) Search(c *gin.Context, query *model.SearchProductQuery) ([]model.ProductResponse, int64, error) {
	products, total, err := s.ProductRepository.Search(s.DB, query)
	if err != nil {
		s.Log.WithError(err).Error("error getting product")
		return nil, 0, util.ErrInternalServer
	}

	responses := make([]model.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *converter.ProductToResponse(&product)
	}

	return responses, total, nil
}

func (s *ProductService) Get(c *gin.Context, id int64) (*model.ProductResponse, error) {
	product := new(entity.Product)
	err := s.ProductRepository.FindById(s.DB, product, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Log.Warnf("Failed to get product from database : %+v", err)
		return nil, util.ErrInternalServer
	}

	if product.ID == 0 {
		s.Log.WithError(err).Error("error getting product")
		return nil, util.ErrProductNotFound
	}

	return converter.ProductToResponse(product), nil
}

func (s *ProductService) Create(c *gin.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	product := &entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
		Price:       request.Price,
		Quantity:    request.Quantity,
	}

	if err := s.ProductRepository.Create(s.DB, product); err != nil {
		s.Log.Warnf("Failed create product to database : %+v", err)
		return nil, util.ErrCreateProduct
	}

	return converter.ProductToResponse(product), nil
}
