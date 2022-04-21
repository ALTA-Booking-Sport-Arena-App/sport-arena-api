package user

import (
	_entities "capstone/entities"
)

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
}