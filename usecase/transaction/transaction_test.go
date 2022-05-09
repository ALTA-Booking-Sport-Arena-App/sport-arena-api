package payment

import (
	_entities "capstone/entities"
	"capstone/usecase/payment"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllHistory(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		data, err := userUseCase.GetAllHistory(1)
		assert.Nil(t, err)
		assert.Equal(t, uint(2000), data[0].TotalPrice)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepositoryError{}, payment.NewService())
		_, err := userUseCase.GetAllHistory(1)
		assert.NotNil(t, err)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		data, err := userUseCase.CreateTransaction(_entities.Payment{})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(2000), data.TotalPrice)
	})

	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		data, err := userUseCase.CreateTransaction(_entities.Payment{TotalPrice: 2000})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(2000), data.TotalPrice)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepositoryError{}, payment.NewService())
		data, err := userUseCase.CreateTransaction(_entities.Payment{})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Payment{}, data)
	})
}

func TestProcessPayment(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		err := userUseCase.ProcessPayment(_entities.TransactionNotificationInput{})
		assert.Nil(t, err)
	})

	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		err := userUseCase.ProcessPayment(_entities.TransactionNotificationInput{PaymentType: "creadit_card", TransactionStatus: "capture", FraudStatus: "accept"})
		assert.Nil(t, err)
	})

	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		err := userUseCase.ProcessPayment(_entities.TransactionNotificationInput{TransactionStatus: "settlement"})
		assert.Nil(t, err)
	})

	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepository{}, payment.NewService())
		err := userUseCase.ProcessPayment(_entities.TransactionNotificationInput{TransactionStatus: "deny"})
		assert.Nil(t, err)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewPaymentUseCase(mockTransactionRepositoryError{}, payment.NewService())
		err := userUseCase.ProcessPayment(_entities.TransactionNotificationInput{})
		assert.NotNil(t, err)
	})
}

// === mock success ===
type mockTransactionRepository struct{}

func (m mockTransactionRepository) GetAllHistory(id int) ([]_entities.Payment, error) {
	return []_entities.Payment{
		{TotalPrice: 2000},
	}, nil
}

func (m mockTransactionRepository) CreateTransaction(booking _entities.Payment) (_entities.Payment, error) {
	return _entities.Payment{
		TotalPrice: 2000, PaymentURL: "payment",
	}, nil
}

func (m mockTransactionRepository) UpdateTransaction(transaction _entities.Payment) (_entities.Payment, error) {
	return _entities.Payment{
		TotalPrice: 2000, Status: "paid", PaymentURL: "payment",
	}, nil
}

func (m mockTransactionRepository) GetByIdTransaction(ID int) (_entities.Payment, error) {
	return _entities.Payment{
		TotalPrice: 2000, Status: "paid",
	}, nil
}

// === mock error ===

type mockTransactionRepositoryError struct{}

func (m mockTransactionRepositoryError) GetAllHistory(id int) ([]_entities.Payment, error) {
	return nil, fmt.Errorf("error get all data facility")
}

func (m mockTransactionRepositoryError) CreateTransaction(booking _entities.Payment) (_entities.Payment, error) {
	return _entities.Payment{}, fmt.Errorf("error get all data facility")
}

func (m mockTransactionRepositoryError) UpdateTransaction(transaction _entities.Payment) (_entities.Payment, error) {
	return _entities.Payment{}, fmt.Errorf("error get all data facility")
}

func (m mockTransactionRepositoryError) GetByIdTransaction(ID int) (_entities.Payment, error) {
	return _entities.Payment{}, fmt.Errorf("error get all data facility")
}
