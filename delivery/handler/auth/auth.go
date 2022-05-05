package auth

import (
	"capstone/delivery/helper"
	"capstone/entities"
	"capstone/usecase/auth"
	"errors"

	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
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
		extract, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("S3CR3T"), nil
		})
		if err != nil {
			return errors.New("error extract token")
		}
		if !extract.Valid {
			return errors.New("invalid")
		}
		claims := extract.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		responseToken := map[string]interface{}{
			"token": token,
			"role":  role,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully logged in", responseToken))
	}
}
