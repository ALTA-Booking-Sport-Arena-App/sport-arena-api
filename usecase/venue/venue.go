package venue

import (
	_entities "capstone/entities"
	_venueRepository "capstone/repository/venue"
)

type VenueUseCase struct {
	venueRepository _venueRepository.VenueRepositoryInterface
}

func NewVenueUseCase(venueRepo _venueRepository.VenueRepositoryInterface) VenueUseCaseInterface {
	return &VenueUseCase{
		venueRepository: venueRepo,
	}
}

func (cuc *VenueUseCase) CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	venue, rows, err := cuc.venueRepository.CreateStep2(request, facility)
	return venue, rows, err
}
