package user

import (
	"capstone/delivery/helper"
	_entities "capstone/entities"
	"errors"

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
	if request.FullName == "" || request.FullName == " " {
		return request, errors.New("can't be empty")
	}
	if request.Email == "" || request.Email == " " {
		return request, errors.New("can't be empty")
	}
	if request.Password == "" || request.Password == " " {
		return request, errors.New("can't be empty")
	}
	if request.PhoneNumber == "" || request.PhoneNumber == " " {
		return request, errors.New("can't be empty")
	}
	if request.Username == "" || request.Username == " " {
		return request, errors.New("can't be empty")
	}

	password, _ := helper.HashPassword(request.Password)
	request.Password = password

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

func (ur *UserRepository) UpdateUserImage(image string, idToken int) (int, error) {
	var users []_entities.User
	tx := ur.DB.Model(&users).Where("id = ?", idToken).Update("image", image)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
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

func (ur *UserRepository) GetUserProfile(id int) (_entities.User, error) {
	var users _entities.User

	yx := ur.DB.Where("id = ?", id).First(&users)

	if yx.Error != nil {
		return users, yx.Error
	}

	return users, nil

}

func (ur *UserRepository) RequestOwner(requestOwner _entities.User) (int, error) {
	tx := ur.DB.Save(&requestOwner)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

// ================= ADMIN SECTIONS =================

func (ur *UserRepository) GetListUsers() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.DB.Not("role = ?", "admin").Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) GetListOwners() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.DB.Preload("Venues").Not("role = ?", "admin").Where("status = ?", "approve").Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) GetListOwnerRequests() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.DB.Not("role = ?", "admin").Where("status = ?", "pending").Or("status = ?", "reject").Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) ApproveOwnerRequest(request _entities.User) error {
	tx := ur.DB.Save(&request)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *UserRepository) RejectOwnerRequest(request _entities.User) error {
	tx := ur.DB.Save(&request)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *UserRepository) UpdateAdmin(id int, password string) error {
	var users _entities.User
	tx := ur.DB.Model(&users).Where("id = ?", id).Update("password", password)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}
