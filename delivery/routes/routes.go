package routes

import (
	_authHandler "capstone/delivery/handler/auth"
	_categoryHandler "capstone/delivery/handler/category"
	_facilityHandler "capstone/delivery/handler/facility"
	_paymentHandler "capstone/delivery/handler/payment"
	_userHandler "capstone/delivery/handler/user"
	_venueHandler "capstone/delivery/handler/venue"
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
	e.GET("/lists/users", uh.GetListUsersHandler(), _middlewares.JWTMiddleware())
	e.GET("/lists/owners", uh.GetLIstOwnersHandler(), _middlewares.JWTMiddleware())
	e.GET("/lists/owners-request", uh.GetListOwnerRequestHandler(), _middlewares.JWTMiddleware())
	e.PUT("/verification/approve", uh.ApproveOwnerRequestHandler(), _middlewares.JWTMiddleware())
	e.PUT("/verification/reject", uh.RejectOwnerRequestHandler(), _middlewares.JWTMiddleware())
	e.PUT("/admin/password", uh.UpdateAdminHandler(), _middlewares.JWTMiddleware())
}

func RegisterCategoryPath(e *echo.Echo, uh _categoryHandler.CategoryHandler) {
	e.GET("/category", uh.GetAllCategoryHandler())
	e.POST("/category", uh.CreateCategoryHandler(), _middlewares.JWTMiddleware())
	e.PUT("/category/:id", uh.UpdateCategoryHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/category/:id", uh.DeleteCategoryHandler(), _middlewares.JWTMiddleware())
}

func RegisterFacilityPath(e *echo.Echo, uh _facilityHandler.FacilityHandler) {
	e.GET("/facility", uh.GetAllFacilityHandler())
	e.POST("/facility", uh.CreateFacilityHandler(), _middlewares.JWTMiddleware())
	e.PUT("/facility/:id", uh.UpdateFacilityHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/facility/:id", uh.DeleteFacilityHandler(), _middlewares.JWTMiddleware())
}

func RegisterVenuePath(e *echo.Echo, uh _venueHandler.VenueHandler) {
	e.POST("/venues/step2", uh.CreateStep2Handler(), _middlewares.JWTMiddleware())
	e.POST("/venues/step1", uh.CreateStep1Handler(), _middlewares.JWTMiddleware())
	e.PUT("/venues/step2/:id", uh.UpdateStep2Handler(), _middlewares.JWTMiddleware())
	e.PUT("/venues/step1/:id", uh.UpdateStep1Handler(), _middlewares.JWTMiddleware())
	e.DELETE("/venues/:id", uh.DeleteVenueHandler(), _middlewares.JWTMiddleware())
	e.GET("/venues", uh.GetAllListHandler())
}

func PaymentArenaPath(e *echo.Echo, ph _paymentHandler.PaymentHandler) {
	e.GET("/histories", ph.GetAllHistoryHandler(), _middlewares.JWTMiddleware())
	e.POST("/booking", ph.CreateFacilityHandler(), _middlewares.JWTMiddleware())
}
