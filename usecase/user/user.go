package user

import (
	"capstone/delivery/helper"
	_entities "capstone/entities"
	_userRepository "capstone/repository/user"
	"errors"

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

func (uuc *UserUseCase) GetUserProfile(id int) (_entities.UserResponse, error) {
	//TODO implement me
	userResponse := _entities.UserResponse{}
	user, _, err := uuc.userRepository.GetUserById(id)

	if err != nil {
		return userResponse, err
	}
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}

func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	password, _ := helper.HashPassword(request.Password)
	request.Password = password
	users, err := uuc.userRepository.CreateUser(request)

	if request.FullName == "" {
		return users, errors.New("Can't be empty")
	}
	if request.Email == "" {
		return users, errors.New("Can't be empty")
	}
	if request.Password == "" {
		return users, errors.New("Can't be empty")
	}
	if request.PhoneNumber == "" {
		return users, errors.New("Can't be empty")
	}
	if request.Username == "" {
		return users, errors.New("Can't be empty")
	}

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
	user.Status = "Pending"
	row, err := uuc.userRepository.RequestOwner(user)
	return row, err
}
