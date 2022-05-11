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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("unauthorized", http.StatusBadRequest))
		}
		// binding data
		var venue _entities.Venue
		errBind := c.Bind(&venue)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind data", http.StatusBadRequest))
		}

		// binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind image", http.StatusBadRequest))
		}
		// check file CheckFileExtension
		_, err_check_extension := image.CheckImageFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file extension", http.StatusBadRequest))
		}
		// check file size
		err_check_size := image.CheckImageFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file size", http.StatusBadRequest))
		}
		fileName := "user_profile_id_" + strconv.Itoa(idToken)
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := image.UploadImage("venues", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file", http.StatusBadRequest))
		}

		venue.UserID = uint(idToken)
		imageURL := theUrl
		data, rows, err := eh.venueUseCase.CreateStep1(venue, imageURL)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to create event", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}
		responseVenue := map[string]interface{}{
			"id": data.ID,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create venue", http.StatusOK, responseVenue))
	}
}

func (uh *VenueHandler) CreateStep2Handler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var param helper.VenueRequestFormat

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed created venue", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Success create venue", http.StatusOK))
	}

}

func (uh *VenueHandler) GetAllListHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//query param for searching venue
		name := c.QueryParam("name")
		category, _ := strconv.Atoi(c.QueryParam("category"))

		getVenues, err := uh.venueUseCase.GetAllList(name, category)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get venue", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success to get venue", http.StatusOK, getVenues))
	}
}

func (uh *VenueHandler) UpdateStep2Handler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var param helper.VenueRequestFormat

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}

		VenueID, _ := strconv.Atoi(c.Param("id"))

		var operationalRequest = []_entities.Step2{}
		for _, v := range param.Day {
			fmt.Println(v)
			request := _entities.Step2{
				VenueID:   uint(VenueID),
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
				VenueID:    uint(VenueID),
				FacilityID: i,
			}
			venuefacility = append(venuefacility, request)
		}

		data, rows, err := uh.venueUseCase.UpdateStep2(uint(VenueID), operationalRequest, venuefacility)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update venue", http.StatusOK, data))
	}

}

func (eh *VenueHandler) UpdateStep1Handler() echo.HandlerFunc {
	return func(c echo.Context) error {

		// binding data
		var venue _entities.Venue
		errBind := c.Bind(&venue)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind data", http.StatusBadRequest))
		}

		id, _ := strconv.Atoi(c.Param("id"))

		_, rows, err := eh.venueUseCase.UpdateStep1(venue, uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed update venue", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success update venue", http.StatusBadRequest))
	}
}

func (eh *VenueHandler) UpdateVenueImageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		// binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to bind image", http.StatusBadRequest))
		}
		// check file CheckFileExtension
		_, err_check_extension := image.CheckImageFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file extension", http.StatusBadRequest))
		}
		// check file size
		err_check_size := image.CheckImageFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error checking file size", http.StatusBadRequest))
		}
		fileName := "update_venue_id_" + strconv.Itoa(id)
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := image.UploadImage("venues", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error to upload file", http.StatusBadRequest))
		}
		rows, err := eh.venueUseCase.UpdateVenueImage(theUrl, uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed update venue", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success update venue", http.StatusOK))
	}
}

func (uh *VenueHandler) DeleteVenueHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		rows, err := uh.venueUseCase.DeleteVenue(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("deleted venue failed", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("deleted venue successfully", http.StatusOK))
	}
}

func (uh *VenueHandler) GetVenueByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		venue, rows, err := uh.venueUseCase.GetVenueById(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to fetch data", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}

		step2 := []map[string]interface{}{}
		for i := 0; i < len(venue.Step2); i++ {
			response := map[string]interface{}{
				"day":        venue.Step2[i].Day,
				"open_hour":  venue.Step2[i].OpenHour,
				"close_hour": venue.Step2[i].CloseHour,
				"price":      venue.Step2[i].Price,
			}
			step2 = append(step2, response)
		}

		facility := []map[string]interface{}{}
		for i := 0; i < len(venue.VenueFacility); i++ {
			response := map[string]interface{}{
				"id":        venue.VenueFacility[i].Facility.ID,
				"name":      venue.VenueFacility[i].Facility.Name,
				"icon_name": venue.VenueFacility[i].Facility.IconName,
			}
			facility = append(facility, response)
		}

		responseVenue := map[string]interface{}{
			"id":                venue.ID,
			"name":              venue.Name,
			"description":       venue.Description,
			"user_id":           venue.UserID,
			"image":             venue.Image,
			"city":              venue.City,
			"address":           venue.Address,
			"operational_hours": step2,
			"facility":          facility,
			"category": map[string]interface{}{
				"id":        venue.Category.ID,
				"name":      venue.Category.Name,
				"icon_name": venue.Category.IconName,
			},
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get detail arena", http.StatusOK, responseVenue))
	}
}
