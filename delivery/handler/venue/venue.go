package venue

import (
	"capstone/delivery/helper"
	_entities "capstone/entities"
	_venueUseCase "capstone/usecase/venue"
	"fmt"
	"net/http"

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

		fmt.Println(venuefacility)

		// var facility = param.FacilityID
		data, rows, err := uh.venueUseCase.CreateStep2(operationalRequest, venuefacility)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed create venue"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create venue", data))
	}

}

func (uh *VenueHandler) GetAllListHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//query param for searching venue
		name := c.QueryParam("name")
		getVenues, err := uh.venueUseCase.GetAllList(name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get venues"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success to get venues", getVenues))
	}
}
