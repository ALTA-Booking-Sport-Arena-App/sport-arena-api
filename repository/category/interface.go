package category

import (
	_entities "capstone/entities"
)

type CategoryRepositoryInterface interface {
	GetAllCategory() ([]_entities.Category, error)
}
