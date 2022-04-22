package user

import (
	_entities "capstone/entities"
)

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, request _entities.User) (_entities.User, int, error)
	GetUserById(id int) (_entities.User, int, error)
}
