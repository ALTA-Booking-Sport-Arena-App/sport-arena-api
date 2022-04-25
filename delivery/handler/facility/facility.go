package facility

import (
	"capstone/delivery/helper"
	_facilityUseCase "capstone/usecase/facility"
	"net/http"

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
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get all facilities"))
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

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all facilities", responseFacilities))
	}
}
