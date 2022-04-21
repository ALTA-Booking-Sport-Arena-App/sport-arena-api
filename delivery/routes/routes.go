package routes

import (
	_authHandler "capstone/delivery/handler/auth"
	_userHandler "capstone/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
}

