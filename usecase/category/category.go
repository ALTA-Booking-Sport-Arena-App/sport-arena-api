package category

import (
	_entities "capstone/entities"
	_categoryRepository "capstone/repository/category"
)

type CategoryUseCase struct {
	categoryRepository _categoryRepository.CategoryRepositoryInterface
}

func NewCategoryUseCase(categoryRepo _categoryRepository.CategoryRepositoryInterface) CategoryUseCaseInterface {
	return &CategoryUseCase{
		categoryRepository: categoryRepo,
	}
}

func (cuc *CategoryUseCase) GetAllCategory() ([]_entities.Category, error) {
	category, err := cuc.categoryRepository.GetAllCategory()
	return category, err
}

func (cuc *CategoryUseCase) CreateCategory(request _entities.Category) (_entities.Category, error) {
	category, err := cuc.categoryRepository.CreateCategory(request)
	return category, err
}
