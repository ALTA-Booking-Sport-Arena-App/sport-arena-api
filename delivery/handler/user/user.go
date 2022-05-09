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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error binding data", http.StatusBadRequest))
		}
		_, err := uh.userUseCase.CreateUser(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Register failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully registered", http.StatusOK))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))

		if idToken != userId {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		err := uh.userUseCase.DeleteUser(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully deleted", http.StatusOK, err))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateRequest _entities.User
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))
		// check authorization
		if idToken != userId {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}
		// binding request data
		err := c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		_, rows, err := uh.userUseCase.UpdateUser(userId, updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("User not found", http.StatusBadRequest))
		}
		users, _, _ := uh.userUseCase.GetUserById(userId)
		responseUser := map[string]interface{}{
			"id":    users.ID,
			"name":  users.FullName,
			"email": users.Email,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success update user data", http.StatusOK, responseUser))
	}
}

func (uh *UserHandler) UpdateUserImageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}
		userId, _ := strconv.Atoi(c.Param("userId"))
		// check authorization
		if idToken != userId {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}
		// binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error to bind image", http.StatusBadRequest))
		}
		// check file CheckFileExtension
		_, err_check_extension := image.CheckImageFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error checking file extension", http.StatusBadRequest))
		}
		// check file size
		err_check_size := image.CheckImageFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error checking file size", http.StatusBadRequest))
		}
		fileName := "user_profile_id_" + strconv.Itoa(idToken)
		// upload the image
		theUrl, err_upload_photo := image.UploadImage("users", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error to upload file", http.StatusBadRequest))
		}
		rows, err := uh.userUseCase.UpdateUserImage(theUrl, idToken)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed edit user profile", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("User not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Success to update user profile", http.StatusOK))
	}
}

func (uh *UserHandler) GetUserProfile() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Token not found", http.StatusBadRequest))
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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Get user profile failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully get user profile", http.StatusOK, responseUser))
	}
}

func (uh *UserHandler) RequestOwnerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}
		// binding data
		var requestOwner _entities.User
		errBind := c.Bind(&requestOwner)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind data", http.StatusBadRequest))
		}
		// binding certiifcate
		fileData, fileInfo, err_binding_certificate := c.Request().FormFile("business_certificate")
		if err_binding_certificate != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind certificate", http.StatusBadRequest))
		}
		// check file CheckFileExtension
		_, err_check_extension := certificate.CheckCertificateFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Please upload pdf file", http.StatusBadRequest))
		}
		// check file size
		err_check_size := certificate.CheckCertificateFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("File is too big", http.StatusBadRequest))
		}
		fileName := "user_owner_certificate_" + strconv.Itoa(idToken)
		// upload the certificate
		certificate, err_upload_certificate := certificate.UploadCertificate("owners", fileName, fileData)
		if err_upload_certificate != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error to upload file", http.StatusBadRequest))
		}
		rows, err := uh.userUseCase.RequestOwner(idToken, certificate, requestOwner)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to request for being owner", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("User not found", http.StatusBadRequest))
		}
		if rows == -1 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Please fill all needed fields", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Success to request for being owner", http.StatusOK))
	}
}

func (uh *UserHandler) GetListUsersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		//call GetListUsers function
		listUsers, err := uh.userUseCase.GetListUsers()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get all users", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all users", http.StatusOK, listUsers))
	}
}

func (uh *UserHandler) GetLIstOwnersHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		//call GetListOwner function
		listOwners, err := uh.userUseCase.GetListOwners()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get all owners", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all owners", http.StatusOK, listOwners))
	}
}

func (uh *UserHandler) GetListOwnerRequestHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		//call GetListOwner function
		listOwnerRequest, err := uh.userUseCase.GetListOwnerRequests()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get all owners request", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get all owners request", http.StatusOK, listOwnerRequest))
	}
}

func (uh *UserHandler) ApproveOwnerRequestHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request _entities.User
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// binding request data
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error(), http.StatusBadRequest))
		}
		if request.Status != "approve" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("verification approved failed", http.StatusBadRequest))
		}
		err := uh.userUseCase.ApproveOwnerRequest(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("verification approved failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("verification approved successfully", http.StatusOK))
	}
}

func (uh *UserHandler) RejectOwnerRequestHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request _entities.User
		// check login status
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// binding request data
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error(), http.StatusBadRequest))
		}
		if request.Status != "reject" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("verification reject failed", http.StatusBadRequest))
		}
		err := uh.userUseCase.ApproveOwnerRequest(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("verification reject failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("verification reject successfully", http.StatusOK))
	}
}

func (uh *UserHandler) UpdateAdminHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request _entities.User
		// check login status
		id, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// check role
		role, errRole := _middlewares.ExtractRole(c)
		if errRole != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		if role != "admin" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error(), http.StatusBadRequest))
		}
		err := uh.userUseCase.UpdateAdmin(id, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("updated password failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("updated password successfully", http.StatusOK))
	}
}
