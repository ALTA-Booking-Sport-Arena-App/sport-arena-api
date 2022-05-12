package category

import (
	_entities "capstone/entities"
)

type CategoryUseCaseInterface interface {
	GetAllCategory() (*_entities.Pagination, error)
	CreateCategory(request _entities.Category) (_entities.Category, error)
	UpdateCategory(id uint, request _entities.Category) (_entities.Category, int, error)
	DeleteCategory(id int) error
}
