package order

import "godz/6-order-api-cart/internal/product"

type OrderCreateRequest struct {
	UserID          string                         `json:"user_id" validate:"required"`
	TotalPrice      float64                        `json:"total_price" validate:"required"`
	Products        []product.ProductCreateRequest `json:"products" validate:"required"`
	ShippingAddress string                         `json:"shipping_address" validate:"required"`
}
