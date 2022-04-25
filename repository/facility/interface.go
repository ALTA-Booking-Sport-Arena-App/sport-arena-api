package facility

import (
	_entities "capstone/entities"
)

type FacilityRepositoryInterface interface {
	GetAllFacility() ([]_entities.Facility, error)
	CreateFacility(request _entities.Facility) (_entities.Facility, error)
}
