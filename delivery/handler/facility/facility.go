package facility

import (
	"capstone/delivery/helper"
	_entities "capstone/entities"
	_facilityUseCase "capstone/usecase/facility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FacilityHandler struct {
	facilityUseCase _facilityUseCase.FacilityUseCaseInterface
}

func NewFacilityHandler(c _facilityUseCase.FacilityUseCaseInterface) FacilityHandler {
	return FacilityHandler{
		facilityUseCase: c,
	}
}

func (uh *FacilityHandler) GetAllFacilityHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		facility, err := uh.facilityUseCase.GetAllFacility()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed get all facilities", http.StatusBadRequest))
		}

		responseFacilities := []map[string]interface{}{}
		for i := 0; i < len(facility); i++ {
			response := map[string]interface{}{
				"ID":        facility[i].ID,
				"name":      facility[i].Name,
				"icon_name": facility[i].IconName,
			}
			responseFacilities = append(responseFacilities, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all facilities", http.StatusOK, responseFacilities))
	}
}

func (uh *FacilityHandler) CreateFacilityHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var param _entities.Facility

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		_, err = uh.facilityUseCase.CreateFacility(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("created facility failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("created facility successfully", http.StatusOK))
	}

}

func (uh *FacilityHandler) UpdateFacilityHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		var param _entities.Facility
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.Bind(&param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		_, rows, err := uh.facilityUseCase.UpdateFacility(uint(id), param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("updated facility failed", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("updated facility successfully", http.StatusOK))
	}
}

func (uh *FacilityHandler) DeleteFacilityHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		err := uh.facilityUseCase.DeleteFacility(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("deleted facility failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("deleted facility successfully", http.StatusOK))
	}
}
