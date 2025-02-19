package product

import (
	"godz/6-order-api-cart/pkg/req"
	"godz/6-order-api-cart/pkg/res"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductRepo *ProductRepo
}

type ProductHandlerDeps struct {
	ProductRepo *ProductRepo
}

func NewProductHandler(mux *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepo: deps.ProductRepo,
	}
	mux.HandleFunc("POST /product", handler.CreateProduct())
	mux.HandleFunc("GET /product", handler.GetProductById())
	mux.HandleFunc("GET /products", handler.GetProducts())
	mux.HandleFunc("PATCH /product/{id}", handler.UpdateProduct())
	mux.HandleFunc("DELETE /product/{id}", handler.DeleteProduct())
}

func (handler *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](w, r)
		if err != nil {
			return
		}
		product := NewProduct(body.Title, body.Price, body.Description, body.ImageURL)

		createdProduct, err := handler.ProductRepo.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JsonResponse(w, createdProduct, http.StatusOK)
	}
}

func (handler *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		product, err := handler.ProductRepo.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JsonResponse(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepo.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Title:       body.Title,
			Price:       body.Price,
			Description: body.Description,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JsonResponse(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.ProductRepo.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JsonResponse(w, true, http.StatusOK)
	}
}

func (handler *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := handler.ProductRepo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JsonResponse(w, products, http.StatusOK)
	}
}
