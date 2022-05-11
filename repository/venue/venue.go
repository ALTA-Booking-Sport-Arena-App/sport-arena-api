package venue

import (
	_entities "capstone/entities"
	"errors"
	"fmt"

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

func (ur *VenueRepository) GetAllList(name string, category int) ([]_entities.Venue, error) {
	var venues []_entities.Venue
	var tx *gorm.DB
	if name != "" && category != 0 {
		name = "%" + name + "%"
		tx = ur.DB.Preload("Step2").Where("name LIKE ?", name).Where("category_id = ?", category).Find(&venues)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if name != "" {
		name = "%" + name + "%"
		tx = ur.DB.Preload("Step2").Where("name LIKE ?", name).Find(&venues)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if category != 0 {
		tx = ur.DB.Preload("Step2").Where("category_id = ?", category).Find(&venues)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = ur.DB.Preload("Step2").Find(&venues)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return venues, nil
}

func (ur *VenueRepository) GetOperational() ([]_entities.Step2, error) {
	var operational []_entities.Step2
	tx := ur.DB.Find(&operational)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return operational, nil
}

func (ur *VenueRepository) UpdateStep2(id int, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	var olddata _entities.Step2

	qx := ur.DB.Model(&_entities.Step2{}).Where("venue_id = ?", id).First(&olddata).Error
	if qx != nil {
		return request, 0, qx
	}

	for key := range request {
		if request[key].OpenHour == "" {
			request[key].OpenHour = olddata.OpenHour
		}
		if request[key].CloseHour == "" {
			request[key].CloseHour = olddata.CloseHour
		}
		if request[key].Price == 0 {
			request[key].Price = olddata.Price
		}
	}

	xx := ur.DB.Unscoped().Where("venue_id = ?", id).Delete(&_entities.Step2{}).Error
	if xx != nil {
		return request, 0, xx
	}

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		fmt.Println("yx", yx.Error)
		return request, 0, yx.Error
	}

	var olddatafacility _entities.VenueFacility

	sx := ur.DB.Model(&_entities.VenueFacility{}).Where("venue_id = ?", id).First(&olddatafacility).Error
	if sx != nil {
		return request, 0, sx
	}

	// for key := range facility {
	// 	if facility[key].FacilityID == 0 {
	// 		facility[key].FacilityID = olddatafacility.FacilityID
	// 	}
	// }

	xx = ur.DB.Unscoped().Where("venue_id = ?", id).Delete(&_entities.VenueFacility{}).Error
	if xx != nil {
		return request, 0, xx
	}

	yx = ur.DB.Save(&facility)
	if yx.Error != nil {
		fmt.Println("yx", yx.Error)
		return request, 0, yx.Error
	}

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

func (ur *VenueRepository) UpdateVenueImage(image string, id uint) (int, error) {
	var venue []_entities.Venue
	yx := ur.DB.Model(&venue).Where("id = ?", id).Update("image", image)
	if yx.Error != nil {
		return 0, yx.Error
	}
	if yx.RowsAffected == 0 {
		return 0, yx.Error
	}
	return int(yx.RowsAffected), nil
}

func (ur *VenueRepository) DeleteVenue(id uint) (int, error) {

	var step2 []_entities.Step2
	ux := ur.DB.Where("venue_id = ?", id).Find(&step2)

	if ux.RowsAffected != 0 {
		yx := ur.DB.Unscoped().Where("venue_id = ?", id).Delete(&_entities.Step2{})
		if yx.Error != nil {
			return 0, yx.Error
		}
		if yx.RowsAffected == 0 {
			return 0, yx.Error
		}
	}

	var facility []_entities.VenueFacility
	px := ur.DB.Where("venue_id = ?", id).Find(&facility)

	if px.RowsAffected != 0 {
		ax := ur.DB.Unscoped().Where("venue_id = ?", id).Delete(&_entities.VenueFacility{})
		if ax.Error != nil {
			return 0, ax.Error
		}
		if ax.RowsAffected == 0 {
			return 0, ax.Error
		}
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

func (ur *VenueRepository) GetCategoryById(id int) ([]_entities.Category, error) {
	var category []_entities.Category
	tx := ur.DB.Where("id = ?", id).Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}
