package venue

import (
	_entities "capstone/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllList(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		data, err := venueUseCase.GetAllList("name", 1)
		assert.Nil(t, err)
		assert.Equal(t, "lapangan", data[0].Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		data, err := venueUseCase.GetAllList("name", 1)
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestCreateStep1(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		data, rows, err := venueUseCase.CreateStep1(_entities.Venue{}, "image")
		assert.Nil(t, err)
		assert.Equal(t, "lapangan", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		data, rows, err := venueUseCase.CreateStep1(_entities.Venue{}, "image")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Equal(t, 1, rows)
	})
}

func TestCreateStep2(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		data, rows, err := venueUseCase.CreateStep2([]_entities.Step2{}, []_entities.VenueFacility{})
		assert.Nil(t, err)
		assert.Equal(t, "", data[0].OpenHour)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		data, rows, err := venueUseCase.CreateStep2([]_entities.Step2{}, []_entities.VenueFacility{})
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Equal(t, 0, rows)
	})
}

func TestUpdateStep1(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		data, rows, err := venueUseCase.UpdateStep1(_entities.Venue{}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "lapangan", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		data, rows, err := venueUseCase.UpdateStep1(_entities.Venue{}, 1)
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
		assert.Equal(t, 1, rows)
	})
}

func TestUpdateStep2(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		data, rows, err := venueUseCase.UpdateStep2(1, []_entities.Step2{}, []_entities.VenueFacility{})
		assert.Nil(t, err)
		assert.Equal(t, "", data[0].OpenHour)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		data, rows, err := venueUseCase.UpdateStep2(0, []_entities.Step2{}, []_entities.VenueFacility{})
		assert.NotNil(t, err)
		assert.Nil(t, data)
		assert.Equal(t, 0, rows)
	})
}

func TestDeleteVenue(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		rows, err := venueUseCase.DeleteVenue(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		rows, err := venueUseCase.DeleteVenue(0)
		assert.NotNil(t, err)
		assert.Equal(t, 1, rows)
	})
}

func TestGetVenueById(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		data, rows, err := venueUseCase.GetVenueById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "lapangan", data.Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		data, rows, err := venueUseCase.GetVenueById(0)
		assert.NotNil(t, err)
		assert.Equal(t, 1, rows)
		assert.Nil(t, nil, data)
	})
}

func TestUpdateVenueImage(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepository{})
		rows, err := venueUseCase.UpdateVenueImage("image", 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		venueUseCase := NewVenueUseCase(mockVenueRepositoryError{})
		rows, err := venueUseCase.UpdateVenueImage("image", 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
	})
}

// === mock success ===
type mockVenueRepository struct{}

func (m mockVenueRepository) GetAllList(name string, category int) ([]_entities.Venue, error) {
	return []_entities.Venue{
		{Name: "lapangan"},
	}, nil
}

func (m mockVenueRepository) CreateStep1(request _entities.Venue, image string) (_entities.Venue, int, error) {
	return _entities.Venue{
		Name: "lapangan",
	}, 1, nil
}

func (m mockVenueRepository) CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	return []_entities.Step2{
		{CloseHour: "10 PM"},
	}, 1, nil
}

func (m mockVenueRepository) UpdateStep2(id int, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	return []_entities.Step2{
		{CloseHour: "10 PM"},
	}, 1, nil
}

func (m mockVenueRepository) UpdateStep1(request _entities.Venue, id uint) (_entities.Venue, int, error) {
	return _entities.Venue{
		Name: "lapangan",
	}, 1, nil
}

func (m mockVenueRepository) DeleteVenue(id uint) (int, error) {
	return 1, nil
}

func (m mockVenueRepository) GetVenueById(id int) (_entities.Venue, int, error) {
	return _entities.Venue{
		Name: "lapangan",
	}, 1, nil
}

func (m mockVenueRepository) UpdateVenueImage(image string, id uint) (int, error) {
	return 1, nil
}

func (m mockVenueRepository) GetOperational() ([]_entities.Step2, error) {
	return []_entities.Step2{
		{Price: 5000},
	}, nil
}

func (m mockVenueRepository) GetCategoryById(id int) ([]_entities.Category, error) {
	return []_entities.Category{
		{Name: "category"},
	}, nil
}

// === mock error ===

type mockVenueRepositoryError struct{}

func (m mockVenueRepositoryError) GetAllList(name string, category int) ([]_entities.Venue, error) {
	return nil, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) CreateStep1(request _entities.Venue, image string) (_entities.Venue, int, error) {
	return _entities.Venue{}, 1, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	return nil, 0, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) UpdateStep1(request _entities.Venue, id uint) (_entities.Venue, int, error) {
	return _entities.Venue{}, 1, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) UpdateStep2(id int, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error) {
	return nil, 0, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) DeleteVenue(id uint) (int, error) {
	return 1, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) UpdateVenueImage(image string, id uint) (int, error) {
	return 0, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) GetOperational() ([]_entities.Step2, error) {
	return nil, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) GetVenueById(id int) (_entities.Venue, int, error) {
	return _entities.Venue{}, 1, fmt.Errorf("error get all data facility")
}

func (m mockVenueRepositoryError) GetCategoryById(id int) ([]_entities.Category, error) {
	return nil, fmt.Errorf("error get all data facility")
}
