package repository

import "github.com/rikughi/ecom-api/internal/entity"

type OrderItemRepository struct {
	Repository[entity.OrderItem]
}

func NewOrderItemRepository() *OrderItemRepository {
	return &OrderItemRepository{}
}
