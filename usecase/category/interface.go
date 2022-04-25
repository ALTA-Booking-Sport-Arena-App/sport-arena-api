package category

import (
	_entities "capstone/entities"
)

type CategoryUseCaseInterface interface {
	GetAllCategory() ([]_entities.Category, error)
}
