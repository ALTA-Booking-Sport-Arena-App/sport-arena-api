package user

import (
	"capstone/delivery/helper"
	_middlewares "capstone/delivery/middlewares"
	_userUseCase "capstone/usecase/user"
	"net/http"
	"strconv"

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

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))

		if idToken != userId {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}

		err := uh.userUseCase.DeleteUser(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully deleted", err))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var param _entities.User
		userId, _ := strconv.Atoi(c.Param("userId"))

		if idToken != userId {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		users, rows, err := uh.userUseCase.UpdateUser(userId, param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseUser := map[string]interface{}{
			"id":    users.ID,
			"name":  users.Fullname,
			"email": users.Email,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data", responseUser))
	}
}
