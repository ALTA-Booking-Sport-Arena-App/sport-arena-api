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
	user, err := uuc.userRepository.GetUserProfile(id)

	if err != nil {
		return userResponse, err
	}

	copier.Copy(&userResponse, &user)

	return userResponse, nil
}

func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	password, err := helper.HashPassword(request.Password)
	request.Password = password
	users, err := uuc.userRepository.CreateUser(request)

	if request.Fullname == "" {
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
	password, _ := helper.HashPassword(request.Password)
	request.Password = password
	user, rows, err := uuc.userRepository.GetUserById(id)
	if err != nil {
		return user, 0, err
	}
	if rows == 0 {
		return user, 0, nil
	}
	if request.Fullname != "" {
		user.Fullname = request.Fullname
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
	if request.Password != "" {
		user.Password = request.Password
	}
	users, rows, err := uuc.userRepository.UpdateUser(user)
	return users, rows, err
}

func (uuc *UserUseCase) UpdateUserImage(image string, idToken int) (int, error) {
	rows, err := uuc.userRepository.UpdateUserImage(image, idToken)
	return rows, err
}
func (uuc *UserUseCase) GetUserById(id int) (_entities.User, int, error) {
	users, rows, err := uuc.userRepository.GetUserById(id)
	return users, rows, err
}
