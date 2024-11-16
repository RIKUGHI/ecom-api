package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/repository"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderService struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	ProductRepository   *repository.ProductRepository
	OrderRepository     *repository.OrderRepository
	OrderItemRepository *repository.OrderItemRepository
}

func NewOrderService(
	DB *gorm.DB,
	Log *logrus.Logger,
	ProductRepository *repository.ProductRepository,
	OrderRepository *repository.OrderRepository,
	OrderItemRepository *repository.OrderItemRepository,
) *OrderService {
	return &OrderService{
		DB:                  DB,
		Log:                 Log,
		ProductRepository:   ProductRepository,
		OrderRepository:     OrderRepository,
		OrderItemRepository: OrderItemRepository,
	}
}

func (s *OrderService) Create(c *gin.Context, request *model.CartCheckoutRequest) (response *model.CheckoutResponse, e error) {
	tx := s.DB.WithContext(c).Begin()
	defer tx.Rollback()

	order := &entity.Order{
		UserID:  1,
		Status:  "pending",
		Address: "some address",
	}

	if err := s.OrderRepository.Create(tx, order); err != nil {
		s.Log.Warnf("Failed create order to database : %+v", err)
		return nil, util.ErrCreateOrder
	}

	// orderItems := make([]*entity.OrderItem, len(request.Items))
	var total float64

	for _, item := range request.Items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("Invalid quantity for product_id: %d", item.ProductID)
		}

		product := new(entity.Product)
		err := s.ProductRepository.FindById(s.DB, product, item.ProductID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.Warnf("Failed to get product from database : %+v", err)
			return nil, util.ErrInternalServer
		}

		if product.ID == 0 {
			s.Log.WithError(err).Error("error getting product")
			return nil, util.ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return nil, fmt.Errorf("Product %s is not available in the quantity requested", product.Name)
		}

		total += product.Price * float64(item.Quantity)

		if err := s.OrderItemRepository.Create(tx, &entity.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}); err != nil {
			s.Log.Warnf("Failed to create order item to database : %+v", err)
			return nil, util.ErrCreateOrderItem
		}

		product.Quantity -= item.Quantity
		if err := s.ProductRepository.Update(tx, product); err != nil {
			s.Log.Warnf("Failed to update product's stock to database : %+v", err)
			return nil, errors.New("Failed to update stock")
		}
	}

	order.Total = total
	if err := s.OrderRepository.Update(tx, order); err != nil {
		s.Log.Warnf("Failed to update order to database : %+v", err)
		return nil, util.ErrCreateOrderItem
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("error creating order")
		return nil, util.ErrInternalServer
	}

	return &model.CheckoutResponse{
		TotalPrice: total,
		OrderID:    order.ID,
	}, nil
}
