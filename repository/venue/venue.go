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

func (ur *VenueRepository) GetAllList(name string, category string) ([]_entities.Venue, error) {
	var venues []_entities.Venue
	var tx *gorm.DB
	if name != "" || category != "" {
		name = "%" + name + "%"
		category = "%" + category + "%"
		tx = ur.DB.Where("name LIKE ? OR category LIKE ?", name, category).Find(&venues)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = ur.DB.Find(&venues)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return venues, nil
}
