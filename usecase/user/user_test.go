package user

import (
	_entities "capstone/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.GetUserById(1)
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.FullName)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.GetUserById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{FullName: "odi"})
		assert.Nil(t, nil, err)
		assert.Equal(t, "haudhi", data.FullName)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.CreateUser(_entities.User{FullName: "odi"})
		assert.NotNil(t, err)
		assert.Equal(t, "", data.FullName)
		assert.Nil(t, nil, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{FullName: "almas"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.FullName)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{FullName: "almas", Password: "123"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.FullName)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{FullName: "almas", Password: "123", Username: "odi"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.FullName)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{FullName: "almas", Email: "odi@mail.com"})
		assert.Nil(t, err)
		assert.Equal(t, "almas", data.FullName)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.UpdateUser(1, _entities.User{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Nil(t, nil, data.FullName)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.DeleteUser(1)
		assert.Nil(t, err)

	})

	t.Run("TestDeleteUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		err := userUseCase.DeleteUser(1)
		assert.NotNil(t, err)

	})
}

func TestGetUserProfile(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.GetUserProfile(1)
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.FullName)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.GetUserProfile(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", data.FullName)
	})
}

func TestUpdateUserImage(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		rows, err := userUseCase.UpdateUserImage("url", 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		rows, err := userUseCase.UpdateUserImage("url", 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
	})
}

func TestRequestOwner(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		rows, err := userUseCase.RequestOwner(1, "certificate", _entities.User{BusinessName: "golf"})
		assert.Nil(t, err)
		assert.Equal(t, -1, rows)
	})

	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		rows, err := userUseCase.RequestOwner(1, "certificate", _entities.User{BusinessName: "golf", BusinessDescription: "golf"})
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		rows, err := userUseCase.RequestOwner(1, "certificate", _entities.User{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
	})
}

func TestGetListUsers(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		_, err := userUseCase.GetListUsers()
		assert.Nil(t, err)
		// assert.Equal(t, "odi", data[0].FullName)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		_, err := userUseCase.GetListUsers()
		assert.NotNil(t, err)
		// assert.Equal(t, "odi", data[0].FullName)
	})
}

func TestGetListOwners(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		_, err := userUseCase.GetListOwners()
		assert.Nil(t, err)
		// assert.Equal(t, "odi", data[0].FullName)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		_, err := userUseCase.GetListOwners()
		assert.NotNil(t, err)
		// assert.Equal(t, "odi", data[0].FullName)
	})
}

func TestApproveOwnerRequest(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.ApproveOwnerRequest(_entities.User{Role: "owner"})
		assert.Nil(t, err)
	})

	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.ApproveOwnerRequest(_entities.User{Role: "owner", Status: "approved"})
		assert.Nil(t, err)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		err := userUseCase.ApproveOwnerRequest(_entities.User{})
		assert.NotNil(t, err)
	})
}

func TestRejectOwnerRequest(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.RejectOwnerRequest(_entities.User{Status: "rejected"})
		assert.Nil(t, err)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		err := userUseCase.RejectOwnerRequest(_entities.User{})
		assert.NotNil(t, err)
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.UpdateAdmin(1, _entities.User{Password: "123"})
		assert.Nil(t, nil, err)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		err := userUseCase.UpdateAdmin(1, _entities.User{})
		assert.NotNil(t, err)
	})
}

func TestGetListOwnerRequest(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		_, err := userUseCase.GetListOwnerRequests()
		assert.Nil(t, err)
		// assert.Equal(t, "odi", data[0].FullName)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		_, err := userUseCase.GetListOwnerRequests()
		assert.NotNil(t, err)
		// assert.Equal(t, "odi", data[0].FullName)
	})
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) GetUserById(id int) (_entities.User, int, error) {
	return _entities.User{
		FullName: "odi", Email: "odi@mail.com", Password: "lalala",
	}, 1, nil
}

func (m mockUserRepository) CreateUser(request _entities.User) (_entities.User, error) {
	return _entities.User{
		FullName: "haudhi", Email: "odi@mail.com", Password: "lalala",
	}, nil
}

func (m mockUserRepository) UpdateUser(request _entities.User) (_entities.User, int, error) {
	return _entities.User{
		FullName: "almas", Email: "odi@mail.com", Password: "lalala",
	}, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) error {
	return nil
}

func (m mockUserRepository) GetUserProfile(id int) (_entities.User, error) {
	return _entities.User{
		FullName: "odi", Email: "odi@mail.com", Password: "lalala",
	}, nil
}

func (m mockUserRepository) UpdateUserImage(image string, idToken int) (int, error) {
	return 1, nil
}

func (m mockUserRepository) RequestOwner(requestOwner _entities.User) (int, error) {
	return 1, nil
}

func (m mockUserRepository) GetListUsers() ([]_entities.User, error) {
	return []_entities.User{
		{FullName: "odi", Email: "odi@mail.com", Password: "lalala"},
	}, nil
}

func (m mockUserRepository) GetListOwners() ([]_entities.User, error) {
	return []_entities.User{
		{FullName: "odi", Email: "odi@mail.com", Password: "lalala"},
	}, nil
}

func (m mockUserRepository) ApproveOwnerRequest(request _entities.User) error {
	return nil
}

func (m mockUserRepository) RejectOwnerRequest(request _entities.User) error {
	return nil
}

func (m mockUserRepository) UpdateAdmin(id int, password string) error {
	return nil
}

func (m mockUserRepository) GetListOwnerRequests() ([]_entities.User, error) {
	return []_entities.User{
		{FullName: "odi", Email: "odi@mail.com", Password: "lalala"},
	}, nil
}

// === mock error ===

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) GetUserById(id int) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error get data user")
}

func (m mockUserRepositoryError) CreateUser(request _entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create data user")
}

func (m mockUserRepositoryError) UpdateUser(request _entities.User) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) DeleteUser(id int) error {
	return fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) GetUserProfile(id int) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create data user")
}

func (m mockUserRepositoryError) UpdateUserImage(image string, idToken int) (int, error) {
	return 0, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) RequestOwner(requestOwner _entities.User) (int, error) {
	return 0, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) GetListUsers() ([]_entities.User, error) {
	return nil, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) GetListOwners() ([]_entities.User, error) {
	return nil, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) ApproveOwnerRequest(request _entities.User) error {
	return fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) RejectOwnerRequest(request _entities.User) error {
	return fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) UpdateAdmin(id int, password string) error {
	return fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) GetListOwnerRequests() ([]_entities.User, error) {
	return nil, fmt.Errorf("error update data user")
}
