package user

import (
	_entities "capstone/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(request _entities.User) (_entities.User, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}

func (ur *UserRepository) DeleteUser(id int) error {

	err := ur.DB.Delete(&_entities.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUser(request _entities.User) (_entities.User, int, error) {
	tx := ur.DB.Save(&request)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}

func (ur *UserRepository) GetUserById(idToken int) (_entities.User, int, error) {
	var users _entities.User
	tx := ur.DB.Where("id = ?", idToken).Find(&users)
	if tx.Error != nil {
		return users, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, 0, nil
	}
	return users, int(tx.RowsAffected), nil
}
