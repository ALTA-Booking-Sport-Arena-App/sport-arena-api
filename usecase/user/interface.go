package user

import (
	_entities "capstone/entities"
)

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, request _entities.User) (_entities.User, int, error)
	GetUserById(id int) (_entities.User, int, error)
	GetUserProfile(id int) (_entities.UserResponse, error)
	UpdateUserImage(image string, idToken int) (int, error)
	RequestOwner(id int, certificate string, requestOwner _entities.User) (int, error)
	GetListUsers() ([]_entities.ListUsersResponse, error)
	GetListOwners() ([]_entities.ListOwnersResponse, error)
	ApproveOwnerRequest(request _entities.User) error
	RejectOwnerRequest(request _entities.User) error
	UpdateAdmin(id int, request _entities.User) error
}
