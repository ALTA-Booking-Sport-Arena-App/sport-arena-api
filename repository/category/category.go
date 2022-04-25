package category

import (
	_entities "capstone/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (cr *CategoryRepository) GetAllCategory() ([]_entities.Category, error) {
	var category []_entities.Category
	tx := cr.DB.Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}

func (ur *CategoryRepository) CreateCategory(request _entities.Category) (_entities.Category, error) {

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}

func (ur *CategoryRepository) UpdateCategory(id uint, request _entities.Category) (_entities.Category, int, error) {
	tx := ur.DB.Model(&_entities.Category{}).Where("id = ?", id).Updates(request)
	if tx.Error != nil {
		return request, 0, tx.Error
	}
	return request, int(tx.RowsAffected), nil
}
