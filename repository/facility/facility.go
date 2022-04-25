package facility

import (
	_entities "capstone/entities"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	DB *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) *FacilityRepository {
	return &FacilityRepository{
		DB: db,
	}
}

func (cr *FacilityRepository) GetAllFacility() ([]_entities.Facility, error) {
	var facility []_entities.Facility
	tx := cr.DB.Find(&facility)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return facility, nil
}
