package entity

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	FirstName string    `gorm:"column:firstName"`
	LastName  string    `gorm:"column:lastName"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime"`
	Token     string    `gorm:"-"`
}

func (u *User) TableName() string {
	return "users"
}
