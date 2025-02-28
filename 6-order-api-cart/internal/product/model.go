package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Price       float64
	Description string
	ImageURL    string
}

func NewProduct(title string, price float64, description, imageURL string) *Product {
	return &Product{
		Title:       title,
		Price:       price,
		Description: description,
		ImageURL:    imageURL,
	}
}
