package venue

import (
	_entities "capstone/entities"
)

type VenueRepositoryInterface interface {
	GetAllList(name string, category int) ([]_entities.Venue, error)
	CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
	CreateStep1(request _entities.Venue, image string) (_entities.Venue, int, error)
	UpdateStep1(request _entities.Venue, id uint) (_entities.Venue, int, error)
	UpdateStep2(id int, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
	DeleteVenue(id uint) (int, error)
	GetVenueById(id int) (_entities.Venue, int, error)
	UpdateVenueImage(image string, id uint) (int, error)
	GetOperational() ([]_entities.Step2, error)
	GetCategoryById(id int) ([]_entities.Category, error)
}
