package auth

import (
	"capstone/delivery/helper"
	"capstone/entities"
	"capstone/usecase/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUseCase auth.AuthUseCaseInterface
}

func NewAuthHandler(auth auth.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: auth,
	}
}

func (ah *AuthHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login entities.User
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error to bind data"))
		}
		token, errorLogin := ah.authUseCase.Login(login.Email, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully logged in", responseToken))
	}
}
