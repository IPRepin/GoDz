package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	title       string
	price       float64
	description string
}

func NewProduct(title string, price float64, description string) *Product {
	return &Product{
		title:       title,
		price:       price,
		description: description,
	}
}
