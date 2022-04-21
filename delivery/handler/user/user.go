package user

import (
	"capstone/delivery/helper"
	_userUseCase "capstone/usecase/user"
	"net/http"

	_entities "capstone/entities"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(u _userUseCase.UserUseCaseInterface) UserHandler {
	return UserHandler{
		userUseCase: u,
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		var param _entities.User

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error binding data"))
		}
		_, err := uh.userUseCase.CreateUser(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Register failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully registered"))
	}
}

