package entity

import "time"

type Order struct {
	ID        int       `gorm:"column:id;primaryKey"`
	UserID    int       `gorm:"column:userId"`
	Total     float64   `gorm:"column:total"`
	Status    string    `gorm:"column:status"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime"`
}

func (o *Order) TableName() string {
	return "orders"
}
