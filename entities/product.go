package entities

import "gorm.io/gorm"

type Product struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Price    float64
	Quantity float64
	gorm.Model
}

func (Product) TableName() string {
	return "product"
}
