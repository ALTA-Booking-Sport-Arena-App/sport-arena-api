package venue

import (
	_entities "capstone/entities"
)

type VenueUseCaseInterface interface {
	CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
}
