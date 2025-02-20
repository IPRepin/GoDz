package order

import (
	"godz/6-order-api-cart/internal/product"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	TotalPrice      float64
	Products        []product.Product `gorm:"many2many:order_products;"`
	ShippingAddress string
}

func NewOrder(
	userId uint,
	totalPrice float64,
	products []product.Product,
	shippingAddress string,
) *Order {
	return &Order{
		UserID:          userId,
		TotalPrice:      totalPrice,
		Products:        products,
		ShippingAddress: shippingAddress,
	}
}
