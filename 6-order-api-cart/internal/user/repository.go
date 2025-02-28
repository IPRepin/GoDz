package user

import "godz/6-order-api-cart/pkg/db"

type UserRepository struct {
	DataBase *db.Db
}

func NewUserRepository(dataBase *db.Db) *UserRepository {
	return &UserRepository{
		DataBase: dataBase,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.DataBase.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.DataBase.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) UpdateSessionId(newSessionId, phone string) (*User, error) {
	result := repo.DataBase.DB.
		Model(&User{}).
		Where("phone = ?", phone).
		Update("session_id", newSessionId)
	if result.Error != nil {
		return nil, result.Error
	}
	var user User
	result = repo.DataBase.DB.
		Where("phone = ?", phone).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
