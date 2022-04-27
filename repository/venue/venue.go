package venue

import (
	_entities "capstone/entities"

	"gorm.io/gorm"
)

type VenueRepository struct {
	DB *gorm.DB
}

func NewVenueRepository(db *gorm.DB) *VenueRepository {
	return &VenueRepository{
		DB: db,
	}
}

func (ur *VenueRepository) CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, 0, yx.Error
	}

	tx := ur.DB.Save(&facility)
	if tx.Error != nil {
		return request, 0, tx.Error
	}

	return request, int(yx.RowsAffected), nil
}
