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
