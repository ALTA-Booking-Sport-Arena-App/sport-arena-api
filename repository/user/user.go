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
