package main

import (
	"capstone/configs"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authHandler "capstone/delivery/handler/auth"
	_paymentHandler "capstone/delivery/handler/payment"
	_authRepository "capstone/repository/auth"
	_paymentRepo "capstone/repository/payment"
	_authUseCase "capstone/usecase/auth"
	_paymentUseCase "capstone/usecase/payment"

	_userHandler "capstone/delivery/handler/user"
	_routes "capstone/delivery/routes"
	_userRepository "capstone/repository/user"
	_userUseCase "capstone/usecase/user"

	_categoryHandler "capstone/delivery/handler/category"
	_categoryRepository "capstone/repository/category"
	_categoryUseCase "capstone/usecase/category"

	_facilityHandler "capstone/delivery/handler/facility"
	_facilityRepository "capstone/repository/facility"
	_facilityUseCase "capstone/usecase/facility"

	_venueHandler "capstone/delivery/handler/venue"
	_venueRepository "capstone/repository/venue"
	_venueUseCase "capstone/usecase/venue"

	_middleware "capstone/delivery/middlewares"
	_utils "capstone/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	categoryRepo := _categoryRepository.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUseCase)

	facilityRepo := _facilityRepository.NewFacilityRepository(db)
	facilityUseCase := _facilityUseCase.NewFacilityUseCase(facilityRepo)
	facilityHandler := _facilityHandler.NewFacilityHandler(facilityUseCase)

	venueRepo := _venueRepository.NewVenueRepository(db)
	venueUseCase := _venueUseCase.NewVenueUseCase(venueRepo)
	venueHandler := _venueHandler.NewVenueHandler(venueUseCase)

	paymentRepo := _paymentRepo.NewPaymentRepository(db)
	paymentUseCase := _paymentUseCase.NewFacilityUseCase(paymentRepo)
	paymentHandler := _paymentHandler.NewFacilityHandler(paymentUseCase)

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
	_routes.RegisterFacilityPath(e, facilityHandler)
	_routes.RegisterVenuePath(e, venueHandler)
	_routes.PaymentArenaPath(e, paymentHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
