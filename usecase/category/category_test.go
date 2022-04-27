package category

import (
	_entities "capstone/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategory(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepository{})
		data, err := categoryUseCase.GetAllCategory()
		assert.Nil(t, err)
		assert.Equal(t, "category 1", data[0].Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepositoryError{})
		data, err := categoryUseCase.GetAllCategory()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("TestCreateSuccess", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepository{})
		data, err := categoryUseCase.CreateCategory(_entities.Category{Name: "category 1"})
		assert.Nil(t, err)
		assert.Equal(t, "category 1", data.Name)
	})

	t.Run("TestCreateError", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepositoryError{})
		data, err := categoryUseCase.CreateCategory(_entities.Category{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("TestCreateSuccess", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepository{})
		data, rows, err := categoryUseCase.UpdateCategory(1, _entities.Category{Name: "category 1"})
		assert.Nil(t, err)
		assert.Equal(t, "category 1", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestCreateError", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepositoryError{})
		data, rows, err := categoryUseCase.UpdateCategory(1, _entities.Category{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Equal(t, 1, rows)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("TestCreateSuccess", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepository{})
		err := categoryUseCase.DeleteCategory(1)
		assert.Nil(t, err)
	})

	t.Run("TestCreateError", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepositoryError{})
		err := categoryUseCase.DeleteCategory(1)
		assert.NotNil(t, err)
	})
}

// === mock success ===
type mockCategoryRepository struct{}

func (m mockCategoryRepository) GetAllCategory() ([]_entities.Category, error) {
	return []_entities.Category{
		{Name: "category 1"},
	}, nil
}

func (m mockCategoryRepository) CreateCategory(request _entities.Category) (_entities.Category, error) {
	return _entities.Category{
		Name: "category 1",
	}, nil
}

func (m mockCategoryRepository) UpdateCategory(id uint, request _entities.Category) (_entities.Category, int, error) {
	return _entities.Category{
		Name: "category 1",
	}, 1, nil
}

func (m mockCategoryRepository) DeleteCategory(id int) error {
	return nil
}

// === mock error ===

type mockCategoryRepositoryError struct{}

func (m mockCategoryRepositoryError) GetAllCategory() ([]_entities.Category, error) {
	return nil, fmt.Errorf("error get all data user")
}

func (m mockCategoryRepositoryError) CreateCategory(request _entities.Category) (_entities.Category, error) {
	return _entities.Category{}, fmt.Errorf("error get all data user")
}

func (m mockCategoryRepositoryError) UpdateCategory(id uint, request _entities.Category) (_entities.Category, int, error) {
	return _entities.Category{}, 1, fmt.Errorf("error get all data user")
}

func (m mockCategoryRepositoryError) DeleteCategory(id int) error {
	return fmt.Errorf("error get all data user")
}
