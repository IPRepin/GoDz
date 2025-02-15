package product

type ProductCreateRequest struct {
	Title       string  `json:"title" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type ProductUpdateRequest struct {
	Title       string  `json:"title" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
