package venue

import (
	_entities "capstone/entities"
	"errors"

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

func (ur *VenueRepository) CreateStep1(request _entities.Venue, image string) (_entities.Venue, int, error) {
	request.Image = image
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, 0, yx.Error
	}
	if yx.RowsAffected == 0 {
		return request, 0, errors.New("unable to save event")
	}
	return request, int(yx.RowsAffected), nil
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

func (ur *VenueRepository) UpdateStep2(id int, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	yx := ur.DB.Model(&[]_entities.Step2{}).Where("venue_id = ?", id).Updates(&request)
	if yx.Error != nil {
		return request, 0, yx.Error
	}

	tx := ur.DB.Model(&[]_entities.VenueFacility{}).Where("venue_id = ?", id).Updates(&facility)
	if tx.Error != nil {
		return request, 0, tx.Error
	}

	// yx := ur.DB.Save(&request)
	// if yx.Error != nil {
	// 	return request, 0, yx.Error
	// }

	// tx := ur.DB.Save(&facility)
	// if tx.Error != nil {
	// 	return request, 0, tx.Error
	// }

	return request, int(yx.RowsAffected), nil
}

func (ur *VenueRepository) UpdateStep1(request _entities.Venue, id uint) (_entities.Venue, int, error) {
	yx := ur.DB.Model(&_entities.Venue{}).Where("id = ?", id).Updates(request)
	if yx.Error != nil {
		return request, 0, yx.Error
	}
	if yx.RowsAffected == 0 {
		return request, 0, yx.Error
	}
	return request, int(yx.RowsAffected), nil
}

func (ur *VenueRepository) DeleteVenue(id uint) (int, error) {

	yx := ur.DB.Unscoped().Where("venue_id = ?", id).Delete(&_entities.Step2{})
	if yx.Error != nil {
		return 0, yx.Error
	}
	if yx.RowsAffected == 0 {
		return 0, yx.Error
	}

	ax := ur.DB.Unscoped().Where("venue_id = ?", id).Delete(&_entities.VenueFacility{})
	if ax.Error != nil {
		return 0, ax.Error
	}
	if ax.RowsAffected == 0 {
		return 0, ax.Error
	}
	tx := ur.DB.Unscoped().Delete(&_entities.Venue{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (ur *VenueRepository) GetVenueById(id int) (_entities.Venue, int, error) {
	var venue _entities.Venue
	tx := ur.DB.Preload("Category").Preload("Step2").Preload("VenueFacility.Facility").Find(&venue, id)
	if tx.Error != nil {
		return venue, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return venue, 0, nil
	}
	return venue, int(tx.RowsAffected), nil
}

func (ur *VenueRepository) GetVenueFacilityById(id int) ([]_entities.VenueFacility, int, error) {
	var venue []_entities.VenueFacility
	tx := ur.DB.Where("venue_id = ?", id).Find(&venue)
	if tx.Error != nil {
		return venue, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return venue, 0, nil
	}
	return venue, int(tx.RowsAffected), nil
}

func (ur *VenueRepository) GetStep2ById(id int) ([]_entities.Step2, int, error) {
	var venue []_entities.Step2
	tx := ur.DB.Where("venue_id = ?", id).Find(&venue)
	if tx.Error != nil {
		return venue, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return venue, 0, nil
	}
	return venue, int(tx.RowsAffected), nil
}
