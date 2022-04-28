package venue

import (
	_entities "capstone/entities"
)

type VenueRepositoryInterface interface {
	GetAllList(name string, category string) ([]_entities.Venue, error)
	CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
}
