package user

import (
	_entities "capstone/entities"
)

type UserRepositoryInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	DeleteUser(id int) error
	UpdateUser(request _entities.User) (_entities.User, int, error)
	GetUserById(id int) (_entities.User, int, error)
	GetUserProfile(id int) (_entities.User, error)
	UpdateUserImage(image string, idToken int) (int, error)
}
