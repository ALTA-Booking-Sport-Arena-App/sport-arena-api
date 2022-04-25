package facility

import (
	_entities "capstone/entities"
)

type FacilityUseCaseInterface interface {
	GetAllFacility() ([]_entities.Facility, error)
	CreateFacility(request _entities.Facility) (_entities.Facility, error)
	UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, int, error)
}
