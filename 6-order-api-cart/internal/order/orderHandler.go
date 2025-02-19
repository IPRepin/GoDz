package order

import (
	"godz/6-order-api-cart/configs"
	"godz/6-order-api-cart/pkg/middleware"
	"net/http"
)

type OrderHandler struct {
	OrderRepo *OrderRepo
}

type OrderHandlerDeps struct {
	OrderRepo *OrderRepo
	Config    *configs.Config
}

func NewOrderHandler(mux *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepo: deps.OrderRepo,
	}
	mux.Handle("POST /orders", middleware.IsAuth(handler.Create(), deps.Config))
	mux.Handle("GET /orders/{id}", middleware.IsAuth(handler.GetById(), deps.Config))
	mux.Handle("GET /my-orders", middleware.IsAuth(handler.GetAllByUserId(), deps.Config))
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OrderHandler) GetAllByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
