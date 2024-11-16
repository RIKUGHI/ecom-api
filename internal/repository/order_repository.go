package repository

import "github.com/rikughi/ecom-api/internal/entity"

type OrderRepository struct {
	Repository[entity.Order]
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}
