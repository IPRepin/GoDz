package order

import (
	"godz/6-order-api-cart/pkg/db"
)

type OrderRepo struct {
	DB *db.Db
}

func NewOrderRepo(db *db.Db) *OrderRepo {
	return &OrderRepo{
		DB: db,
	}
}

func (repo *OrderRepo) Create(order *Order) (*Order, error) {
	result := repo.DB.Create(order)
	if result.Error != nil {
		return order, result.Error
	}
	return order, nil
}

func (repo *OrderRepo) GetById(id uint) (*Order, error) {
	var order Order
	result := repo.DB.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepo) GetAllByUserId(userId uint) ([]Order, error) {
	var orders []Order
	result := repo.DB.Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
