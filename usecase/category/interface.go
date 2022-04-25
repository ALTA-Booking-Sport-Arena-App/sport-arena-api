package category

import (
	_entities "capstone/entities"
)

type CategoryUseCaseInterface interface {
	GetAllCategory() ([]_entities.Category, error)
	CreateCategory(request _entities.Category) (_entities.Category, error)
}
