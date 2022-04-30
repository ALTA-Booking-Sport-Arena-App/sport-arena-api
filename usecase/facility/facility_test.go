package facility

import (
	_entities "capstone/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFacility(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepository{})
		data, err := facilityUseCase.GetAllFacility()
		assert.Nil(t, err)
		assert.Equal(t, "toilet", data[0].Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepositoryError{})
		data, err := facilityUseCase.GetAllFacility()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestCreateFacility(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepository{})
		data, err := facilityUseCase.CreateFacility(_entities.Facility{})
		assert.Nil(t, err)
		assert.Equal(t, "toilet", data.Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepositoryError{})
		data, err := facilityUseCase.CreateFacility(_entities.Facility{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

func TestUpdateFacility(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepository{})
		data, rows, err := facilityUseCase.UpdateFacility(1, _entities.Facility{})
		assert.Nil(t, err)
		assert.Equal(t, "toilet", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepositoryError{})
		data, rows, err := facilityUseCase.UpdateFacility(1, _entities.Facility{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Equal(t, 1, rows)
	})
}

func TestDeleteFacility(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepository{})
		err := facilityUseCase.DeleteFacility(1)
		assert.Nil(t, err)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockFacilityRepositoryError{})
		err := facilityUseCase.DeleteFacility(1)
		assert.NotNil(t, err)
	})
}

// === mock success ===
type mockFacilityRepository struct{}

func (m mockFacilityRepository) GetAllFacility() ([]_entities.Facility, error) {
	return []_entities.Facility{
		{Name: "toilet"},
	}, nil
}

func (m mockFacilityRepository) CreateFacility(request _entities.Facility) (_entities.Facility, error) {
	return _entities.Facility{
		Name: "toilet",
	}, nil
}

func (m mockFacilityRepository) UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, int, error) {
	return _entities.Facility{
		Name: "toilet",
	}, 1, nil
}

func (m mockFacilityRepository) DeleteFacility(id int) error {
	return nil
}

// === mock error ===

type mockFacilityRepositoryError struct{}

func (m mockFacilityRepositoryError) GetAllFacility() ([]_entities.Facility, error) {
	return nil, fmt.Errorf("error get all data facility")
}

func (m mockFacilityRepositoryError) CreateFacility(request _entities.Facility) (_entities.Facility, error) {
	return _entities.Facility{}, fmt.Errorf("error get all data facility")
}

func (m mockFacilityRepositoryError) UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, int, error) {
	return _entities.Facility{}, 1, fmt.Errorf("error get all data facility")
}

func (m mockFacilityRepositoryError) DeleteFacility(id int) error {
	return fmt.Errorf("error get all data facility")
}
