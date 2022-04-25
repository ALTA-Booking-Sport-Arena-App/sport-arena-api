package routes

import (
	_authHandler "capstone/delivery/handler/auth"
	_categoryHandler "capstone/delivery/handler/category"
	_userHandler "capstone/delivery/handler/user"
	_middlewares "capstone/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.DELETE("/users/:userId", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:userId", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/image/:userId", uh.UpdateUserImageHandler(), _middlewares.JWTMiddleware())
	e.GET("/users/profile", uh.GetUserProfile(), _middlewares.JWTMiddleware())
	e.PUT("/owners/request", uh.RequestOwnerHandler(), _middlewares.JWTMiddleware())
}

func RegisterCategoryPath(e *echo.Echo, uh _categoryHandler.CategoryHandler) {
	e.GET("/category", uh.GetAllCategoryHandler())
}
