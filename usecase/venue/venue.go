package venue

import (
	_entities "capstone/entities"
	_venueRepository "capstone/repository/venue"
	"fmt"

	"github.com/jinzhu/copier"
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

func (cuc *VenueUseCase) CreateStep1(request _entities.Venue, image string) (_entities.Venue, int, error) {
	venue, rows, err := cuc.venueRepository.CreateStep1(request, image)
	return venue, rows, err
}

func (cuc *VenueUseCase) GetAllList(name string, category string) ([]_entities.GetVenuesResponse, error) {
	var getVenues []_entities.GetVenuesResponse
	venues, err := cuc.venueRepository.GetAllList(name, category)
	if err != nil {
		return getVenues, err
	}
	copier.Copy(&getVenues, &venues)
	return getVenues, err
}

func (cuc *VenueUseCase) UpdateStep2(VenueID uint, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	// venueFind, rows, err := cuc.venueRepository.GetVenueFacilityById(int(VenueID))
	// facilityFind, rows, err := cuc.venueRepository.GetStep2ById(int(VenueID))
	venue, rows, err := cuc.venueRepository.UpdateStep2(int(VenueID), request, facility)
	fmt.Println(venue)
	return venue, rows, err
}

func (cuc *VenueUseCase) UpdateStep1(request _entities.Venue, id uint) (_entities.Venue, int, error) {
	venue, rows, err := cuc.venueRepository.UpdateStep1(request, id)
	return venue, rows, err
}

func (cuc *VenueUseCase) DeleteVenue(id uint) (int, error) {
	rows, err := cuc.venueRepository.DeleteVenue(id)
	return rows, err
}

func (cuc *VenueUseCase) GetVenueById(id int) (_entities.Venue, int, error) {
	venue, rows, err := cuc.venueRepository.GetVenueById(id)
	return venue, rows, err
}
