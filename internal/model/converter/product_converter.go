package converter

import (
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
)

func ProductToResponse(product *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
	}
}
