package main

import (
	"capstone/configs"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authHandler "capstone/delivery/handler/auth"
	_authRepository "capstone/repository/auth"
	_authUseCase "capstone/usecase/auth"

	_userHandler "capstone/delivery/handler/user"
	_routes "capstone/delivery/routes"
	_userRepository "capstone/repository/user"
	_userUseCase "capstone/usecase/user"

	_categoryHandler "capstone/delivery/handler/category"
	_categoryRepository "capstone/repository/category"
	_categoryUseCase "capstone/usecase/category"

	_middleware "capstone/delivery/middlewares"
	_utils "capstone/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	categoryRepo := _categoryRepository.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(_middleware.CustomLogger())

	_routes.RegisterUserPath(e, userHandler)
	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterCategoryPath(e, categoryHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
