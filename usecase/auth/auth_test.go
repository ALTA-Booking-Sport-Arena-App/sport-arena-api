package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepository{})
		data, err := authUseCase.Login("", "")
		assert.Nil(t, err)
		assert.Equal(t, "odi@mail.com", data)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepositoryError{})
		data, err := authUseCase.Login("", "")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}



// === mock success ===
type mockAuthRepository struct{}

func (m mockAuthRepository) Login(email string, password string) (string, error) {
	return "odi@mail.com", nil
}


// === mock error ===

type mockAuthRepositoryError struct{}

func (m mockAuthRepositoryError) Login(email string, password string) (string, error) {
	return "", fmt.Errorf("error")
}

