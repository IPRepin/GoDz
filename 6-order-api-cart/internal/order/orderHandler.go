package order

import (
	"godz/6-order-api-cart/configs"
	"godz/6-order-api-cart/internal/product"
	"godz/6-order-api-cart/pkg/middleware"
	"godz/6-order-api-cart/pkg/req"
	"godz/6-order-api-cart/pkg/res"
	"net/http"
	"strconv"
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
	mux.Handle("POST /orders", middleware.IsAuth(handler.CreateOrder(), deps.Config))
	mux.Handle("GET /orders/{id}", middleware.IsAuth(handler.GetById(), deps.Config))
	mux.Handle("GET /my-orders", middleware.IsAuth(handler.GetAllByUserId(), deps.Config))
}

func (handler *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[OrderCreateRequest](w, r)
		if err != nil {
			return
		}
		userIDUint := ConvertUserId(body.UserID, w)

		products := product.ConvertProductRequests(body.Products)
		totalPrice := CalculateTotalPrice(products)

		order := NewOrder(userIDUint, totalPrice, products, body.ShippingAddress)

		createdOrder := handler.OrderRepo.DB.Create(order)
		if err != nil {
			res.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
			return
		}

		res.JsonResponse(w, createdOrder, http.StatusCreated)
	}
}

func (handler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idUint, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
		if err != nil {
			res.JsonResponse(w, map[string]string{"error": "некорректный формат id заказа"},
				http.StatusBadRequest)
			return
		}

		order, err := handler.OrderRepo.GetById(uint(idUint))
		if err != nil {
			res.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
			return
		}
		res.JsonResponse(w, order, http.StatusOK)
	}
}

func (handler *OrderHandler) GetAllByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDValue := r.Context().Value("user_id")
		if userIDValue == nil {
			res.JsonResponse(w, map[string]string{"error": "идентификатор пользователя не найден"},
				http.StatusUnauthorized)
			return
		}

		userID, ok := userIDValue.(uint)
		if !ok {
			res.JsonResponse(w, map[string]string{"error": "неверный формат идентификатора пользователя"},
				http.StatusInternalServerError)
			return
		}

		orders, err := handler.OrderRepo.GetAllByUserId(userID)
		if err != nil {
			res.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
			return
		}

		res.JsonResponse(w, orders, http.StatusOK)
	}
}
