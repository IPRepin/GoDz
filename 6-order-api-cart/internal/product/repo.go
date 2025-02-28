package product

import (
	"godz/6-order-api-cart/pkg/db"
)

type ProductRepo struct {
	DB *db.Db
}

func NewProductRepo(db *db.Db) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (repo *ProductRepo) Create(product *Product) (*Product, error) {
	result := repo.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepo) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepo) Update(product *Product) (*Product, error) {
	result := repo.DB.Save(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepo) Delete(id uint) error {
	result := repo.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepo) GetAll() ([]Product, error) {
	var products []Product
	result := repo.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
