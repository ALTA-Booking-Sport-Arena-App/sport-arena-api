package user

import (
	"capstone/delivery/certificate"
	"capstone/delivery/helper"
	"capstone/delivery/image"
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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error binding data"))
		}
		_, err := uh.userUseCase.CreateUser(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Register failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully registered"))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))

		if idToken != userId {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		err := uh.userUseCase.DeleteUser(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed delete user"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete user"))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var updateRequest _entities.User
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))
		// check authorization
		if idToken != userId {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding request data
		err := c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed edit user profile"))
		}
		_, rows, err := uh.userUseCase.UpdateUser(userId, updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed edit user profile"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("user not found"))
		}
		users, _, _ := uh.userUseCase.GetUserById(userId)
		responseUser := map[string]interface{}{
			"id":    users.ID,
			"name":  users.FullName,
			"email": users.Email,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success edit user profile", responseUser))
	}
}

func (uh *UserHandler) UpdateUserImageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		userId, _ := strconv.Atoi(c.Param("userId"))
		// check authorization
		if idToken != userId {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind image"))
		}
		// check file CheckFileExtension
		_, err_check_extension := image.CheckImageFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file extension"))
		}
		// check file size
		err_check_size := image.CheckImageFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file size"))
		}
		fileName := "user_profile_id_" + strconv.Itoa(idToken)
		// upload the image
		theUrl, err_upload_photo := image.UploadImage("users", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file"))
		}
		rows, err := uh.userUseCase.UpdateUserImage(theUrl, idToken)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed edit user profile"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("user not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success edit user profile"))
	}
}

func (uh *UserHandler) GetUserProfile() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Token not found"))
		}

		userProfile, err := uh.userUseCase.GetUserProfile(id)

		responseUser := map[string]interface{}{
			"id":                    userProfile.ID,
			"fullname":              userProfile.FullName,
			"username":              userProfile.Username,
			"role":                  userProfile.Role,
			"status":                userProfile.Status,
			"email":                 userProfile.Email,
			"image":                 userProfile.Image,
			"phone_number":          userProfile.PhoneNumber,
			"bussiness_name":        userProfile.BusinessName,
			"bussiness_description": userProfile.BusinessDescription,
			"bussiness_certificate": userProfile.BusinessCertificate,
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed get user profile"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get user profile", responseUser))
	}
}

func (uh *UserHandler) RequestOwnerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding data
		var requestOwner _entities.User
		errBind := c.Bind(&requestOwner)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind data"))
		}
		// binding certiifcate
		fileData, fileInfo, err_binding_certificate := c.Request().FormFile("business_certificate")
		if err_binding_certificate != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind certificate"))
		}
		// check file CheckFileExtension
		_, err_check_extension := certificate.CheckCertificateFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("please upload pdf file"))
		}
		// check file size
		err_check_size := certificate.CheckCertificateFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file is too big"))
		}
		fileName := "user_owner_certificate_" + strconv.Itoa(idToken)
		// upload the certificate
		certificate, err_upload_certificate := certificate.UploadCertificate("owners", fileName, fileData)
		if err_upload_certificate != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file"))
		}
		rows, err := uh.userUseCase.RequestOwner(idToken, certificate, requestOwner)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to request for being owner"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("user not found"))
		}
		if rows == -1 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("please fill all needed fields"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to request for being owner"))
	}
}

func (uh *UserHandler) GetListUsersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		//call GetListUsers function
		listUsers, err := uh.userUseCase.GetListUsers()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get all users"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all users", listUsers))
	}
}

func (uh *UserHandler) GetLIstOwnersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		//call GetListOwner function
		listOwners, err := uh.userUseCase.GetListOwners()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get all owners"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all owners", listOwners))
	}
}

func (uh *UserHandler) ApproveOwnerRequestHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request _entities.User
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding request data
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error()))
		}
		err := uh.userUseCase.ApproveOwnerRequest(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("verification approved failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("verification approved successfully"))
	}
}

func (uh *UserHandler) RejectOwnerRequestHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request _entities.User
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding request data
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error()))
		}
		err := uh.userUseCase.ApproveOwnerRequest(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("verification reject failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("verification reject successfully"))
	}
}

func (uh *UserHandler) UpdateAdminHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request _entities.User
		// check login status
		id, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error()))
		}
		err := uh.userUseCase.UpdateAdmin(id, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("updated password failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("updated password successfully"))
	}
}
