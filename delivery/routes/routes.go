package routes

import (
	_userHandler "capstone/delivery/handler/user"

	"github.com/labstack/echo/v4"
)



func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
}

