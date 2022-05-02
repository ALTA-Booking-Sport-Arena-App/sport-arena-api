package payment

import (
	_entities "capstone/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllHistory(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockPaymentRepository{})
		data, err := facilityUseCase.GetAllHistory(1)
		assert.Nil(t, err)
		assert.Equal(t, "paid", data[0].Status)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockPaymentRepositoryError{})
		data, err := facilityUseCase.GetAllHistory(0)
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestCreateBooking(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockPaymentRepository{})
		data, err := facilityUseCase.CreateBooking(_entities.Payment{})
		assert.Nil(t, err)
		assert.Equal(t, "paid", data.Status)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		facilityUseCase := NewFacilityUseCase(mockPaymentRepositoryError{})
		data, err := facilityUseCase.CreateBooking(_entities.Payment{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

// === mock success ===
type mockPaymentRepository struct{}

func (m mockPaymentRepository) GetAllHistory(id int) ([]_entities.Payment, error) {
	return []_entities.Payment{
		{Status: "paid"},
	}, nil
}

func (m mockPaymentRepository) CreateBooking(booking _entities.Payment) (_entities.Payment, error) {
	return _entities.Payment{
		Status: "paid",
	}, nil
}

// === mock error ===

type mockPaymentRepositoryError struct{}

func (m mockPaymentRepositoryError) GetAllHistory(id int) ([]_entities.Payment, error) {
	return nil, fmt.Errorf("error get all data")
}

func (m mockPaymentRepositoryError) CreateBooking(booking _entities.Payment) (_entities.Payment, error) {
	return _entities.Payment{}, fmt.Errorf("error get all data")
}
