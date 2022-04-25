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

func (ur *FacilityRepository) CreateFacility(request _entities.Facility) (_entities.Facility, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}

func (ur *FacilityRepository) UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, int, error) {
	tx := ur.DB.Model(&_entities.Facility{}).Where("id = ?", id).Updates(request)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}

func (ur *FacilityRepository) DeleteFacility(id int) error {

	err := ur.DB.Unscoped().Delete(&_entities.Facility{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
