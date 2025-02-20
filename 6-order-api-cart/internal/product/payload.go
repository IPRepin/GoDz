package product

type ProductCreateRequest struct {
	Title       string  `json:"title" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
	ImageURL    string  `json:"image_url" validate:"required"`
}

type ProductUpdateRequest struct {
	Title       string  `json:"title" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
	ImageURL    string  `json:"image_url" validate:"required"`
}

func ConvertProductRequests(requests []ProductCreateRequest) []Product {
	var products []Product
	for _, req := range requests {
		prod := Product{
			Title:       req.Title,
			Price:       req.Price,
			Description: req.Description,
			ImageURL:    req.ImageURL,
		}
		products = append(products, prod)
	}
	return products
}
