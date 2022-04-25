package facility

import (
	_entities "capstone/entities"
)

type FacilityUseCaseInterface interface {
	GetAllFacility() ([]_entities.Facility, error)
}
