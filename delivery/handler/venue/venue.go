package venue

import (
	"capstone/delivery/helper"
	"capstone/delivery/image"
	_middlewares "capstone/delivery/middlewares"
	_entities "capstone/entities"
	_venueUseCase "capstone/usecase/venue"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VenueHandler struct {
	venueUseCase _venueUseCase.VenueUseCaseInterface
}

func NewVenueHandler(c _venueUseCase.VenueUseCaseInterface) VenueHandler {
	return VenueHandler{
		venueUseCase: c,
	}
}

func (eh *VenueHandler) CreateStep1Handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		// binding data
		var venue _entities.Venue
		errBind := c.Bind(&venue)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("error to bind data"))
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
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := image.UploadImage("venues", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file"))
		}
		// create event
		venue.UserID = uint(idToken)
		imageURL := theUrl
		_, rows, err := eh.venueUseCase.CreateStep1(venue, imageURL)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to create event"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success to create event"))
	}
}

func (uh *VenueHandler) CreateStep2Handler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var param helper.VenueRequestFormat

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		var operationalRequest = []_entities.Step2{}
		for _, v := range param.Day {
			fmt.Println(v)
			request := _entities.Step2{
				VenueID:   param.VenueID,
				OpenHour:  param.OpenHour,
				CloseHour: param.CloseHour,
				Price:     param.Price,
				Day:       v,
			}
			operationalRequest = append(operationalRequest, request)
		}

		var venuefacility = []_entities.VenueFacility{}
		for _, i := range param.FacilityID {
			fmt.Println(i)
			request := _entities.VenueFacility{
				VenueID:    param.VenueID,
				FacilityID: i,
			}
			venuefacility = append(venuefacility, request)

		}

		_, rows, err := uh.venueUseCase.CreateStep2(operationalRequest, venuefacility)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed create venue"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create venue"))
	}

}

func (uh *VenueHandler) GetAllListHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//query param for searching venue
		name := c.QueryParam("name")
		category := c.QueryParam("category")
		getVenues, err := uh.venueUseCase.GetAllList(name, category)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get venues"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get venues", getVenues))
	}
}
