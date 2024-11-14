package entity

import "time"

type Product struct {
	ID          int       `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Image       string    `gorm:"column:image"`
	Price       float64   `gorm:"column:price"`
	Quantity    int       `gorm:"column:quantity"`
	CreatedAt   time.Time `gorm:"column:createdAt;autoCreateTime"`
}

func (u *Product) TableName() string {
	return "products"
}
