package venue

import (
	_entities "capstone/entities"
)

type VenueRepositoryInterface interface {
	CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
}
