package category

import (
	_entities "capstone/entities"
	"math"

	"github.com/jinzhu/copier"
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

func paginate(value interface{}, pagination *_entities.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (cr *CategoryRepository) GetAllCategory() (*_entities.Pagination, error) {
	var category []*_entities.Category
	var categoryResponse []*_entities.CategoryResponse
	var paginationEntity _entities.Pagination

	tx := cr.DB.Scopes(paginate(category, &paginationEntity, cr.DB)).Find(&category)
	copier.Copy(&categoryResponse, &category)

	paginationEntity.DataRow = categoryResponse

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &paginationEntity, nil
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

func (ur *CategoryRepository) DeleteCategory(id int) error {

	err := ur.DB.Unscoped().Delete(&_entities.Category{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
