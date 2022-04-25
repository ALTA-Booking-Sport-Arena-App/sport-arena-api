package facility

import (
	_entities "capstone/entities"
)

type FacilityRepositoryInterface interface {
	GetAllFacility() ([]_entities.Facility, error)
}
