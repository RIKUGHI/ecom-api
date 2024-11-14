package repository

import (
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) Search(db *gorm.DB, query *model.SearchProductQuery) ([]entity.Product, int64, error) {
	var products []entity.Product

	if err := db.Scopes(r.products(query)).Offset((query.Page - 1) * query.Size).Limit(query.Size).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Product{}).Scopes(r.products(query)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *ProductRepository) products(query *model.SearchProductQuery) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if name := query.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}

		return tx
	}
}
