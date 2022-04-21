package user

import (
	_entities "capstone/entities"
)

type UserRepositoryInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	DeleteUser(id int) error
}
