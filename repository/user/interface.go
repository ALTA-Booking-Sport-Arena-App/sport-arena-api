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
	RequestOwner(requestOwner _entities.User) (int, error)
	GetListUsers() ([]_entities.User, error)
	GetListOwners() ([]_entities.User, error)
	ApproveOwnerRequest(request _entities.User) error
	RejectOwnerRequest(request _entities.User) error
}
