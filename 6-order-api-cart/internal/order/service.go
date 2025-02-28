package order

import (
	"godz/6-order-api-cart/internal/product"
	"net/http"
	"strconv"
)

// CalculateTotalPrice суммирует цены всех товаров и возвращает итоговую сумму.
func CalculateTotalPrice(products []product.Product) float64 {
	total := 0.0
	for _, p := range products {
		total += p.Price
	}
	return total
}

func ConvertUserId(userId string, w http.ResponseWriter) uint {
	userIDUint, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return uint(userIDUint)
}
