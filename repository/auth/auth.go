package auth

import (
	"capstone/delivery/helper"
	"capstone/delivery/middlewares"
	"capstone/entities"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) Login(email string, password string) (string, error) {
	var user entities.User
	tx := ar.database.Where("email = ?", email).Find(&user)
	if tx.Error != nil {
		return "failed", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}

	if !helper.CheckPassHash(password, user.Password) {
		return "password incorrect", errors.New("password incorrect")
	}

	token, err := middlewares.CreateToken(int(user.ID), user.FullName, user.Role)
	if err != nil {
		return "create token failed", err
	}
	return token, nil
}
