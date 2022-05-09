package payment

import (
	_entities "capstone/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentURL(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewService()
		data, err := userUseCase.GetPaymentURL(_entities.Payment{}, _entities.User{})
		assert.Nil(t, nil, err)
		assert.Equal(t, "", data)
	})
}
