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

func (order *Order) NewOrder() *Order {
	return &Order{
		UserID:          order.UserID,
		TotalPrice:      order.TotalPrice,
		Products:        make([]product.Product, 0),
		ShippingAddress: order.ShippingAddress,
	}
}
