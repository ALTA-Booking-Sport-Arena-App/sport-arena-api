package user

import (
	"capstone/delivery/helper"
	_entities "capstone/entities"
	_userRepository "capstone/repository/user"
	"fmt"

	"github.com/jinzhu/copier"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetUserProfile(id int) (_entities.User, error) {
	//TODO implement me

	user, _, err := uuc.userRepository.GetUserById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.CreateUser(request)
	return users, err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.userRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) UpdateUser(id int, request _entities.User) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return user, 0, err
	}
	if rows == 0 {
		return user, 0, err
	}
	if request.Password == "" && request.FullName == "" && request.Username == "" && request.Email == "" && request.PhoneNumber == "" {
		return user, -1, err
	}
	if request.Password != "" {
		password, _ := helper.HashPassword(request.Password)
		user.Password = password
	}
	if request.FullName != "" {
		user.FullName = request.FullName
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.PhoneNumber != "" {
		user.PhoneNumber = request.PhoneNumber
	}
	data, rows, err := uuc.userRepository.UpdateUser(user)
	return data, rows, err
}

func (uuc *UserUseCase) UpdateUserImage(image string, idToken int) (int, error) {
	rows, err := uuc.userRepository.UpdateUserImage(image, idToken)
	return rows, err
}

func (uuc *UserUseCase) GetUserById(id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetUserById(id)
	return user, rows, err
}

func (uuc *UserUseCase) RequestOwner(id int, certificate string, requestOwner _entities.User) (int, error) {
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, err
	}
	if requestOwner.BusinessName != "" {
		user.BusinessName = requestOwner.BusinessName
	} else if requestOwner.BusinessName == "" {
		return -1, err
	}
	if requestOwner.BusinessDescription != "" {
		user.BusinessDescription = requestOwner.BusinessDescription
	} else if requestOwner.BusinessDescription == "" {
		return -1, err
	}
	user.BusinessCertificate = certificate
	user.Status = "pending"
	row, err := uuc.userRepository.RequestOwner(user)
	return row, err
}

func (uuc *UserUseCase) GetListUsers() ([]_entities.ListUsersResponse, error) {
	listUsersResponse := []_entities.ListUsersResponse{}
	users, err := uuc.userRepository.GetListUsers()

	if err != nil {
		return listUsersResponse, err
	}
	copier.Copy(&listUsersResponse, &users)
	return listUsersResponse, nil
}

func (uuc *UserUseCase) GetListOwners() ([]_entities.ListOwnersResponse, error) {
	listOwnersResponse := []_entities.ListOwnersResponse{}
	owners, err := uuc.userRepository.GetListOwners()

	if err != nil {
		return listOwnersResponse, err
	}
	copier.Copy(&listOwnersResponse, &owners)
	return listOwnersResponse, nil
}

func (uuc *UserUseCase) GetListOwnerRequests() ([]_entities.ListOwnerRequestResponse, error) {
	listOwnerRequestResponse := []_entities.ListOwnerRequestResponse{}
	owners, err := uuc.userRepository.GetListOwnerRequests()

	if err != nil {
		return listOwnerRequestResponse, err
	}
	copier.Copy(&listOwnerRequestResponse, &owners)
	return listOwnerRequestResponse, nil
}

func (uuc *UserUseCase) ApproveOwnerRequest(request _entities.User) error {
	id := int(request.ID)
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return err
	}
	if request.Role != "" {
		user.Role = request.Role
	}
	if request.Status != "" {
		user.Status = request.Status
	}
	err1 := uuc.userRepository.ApproveOwnerRequest(user)
	if err1 != nil {
		return err1
	}
	return nil
}

func (uuc *UserUseCase) RejectOwnerRequest(request _entities.User) error {
	id := int(request.ID)
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return err
	}
	if request.Status != "" {
		user.Status = request.Status
	}
	err1 := uuc.userRepository.ApproveOwnerRequest(user)
	if err1 != nil {
		return err1
	}
	return nil
}

func (uuc *UserUseCase) UpdateAdmin(id int, request _entities.User) error {
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return err
	}
	if helper.CheckPassHash(request.Password, user.Password) {
		return fmt.Errorf("invalid password")
	}
	if request.Password != "" {
		password, _ := helper.HashPassword(request.Password)
		user.Password = password
	} else if request.Password == "" {
		return fmt.Errorf("invalid password")
	}
	err1 := uuc.userRepository.UpdateAdmin(id, user.Password)
	if err1 != nil {
		return err1
	}
	return nil
}
