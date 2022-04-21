package main

import (
	"capstone/configs"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_userHandler "capstone/delivery/handler/user"
	_routes "capstone/delivery/routes"
	_userRepository "capstone/repository/user"
	_userUseCase "capstone/usecase/user"

	_middleware "capstone/delivery/middlewares"
	_utils "capstone/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	
	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},

	}))
	e.Use(_middleware.CustomLogger())

	_routes.RegisterUserPath(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}