package payment

import (
	_entities "capstone/entities"
	"fmt"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		database: db,
	}
}

func (pr *PaymentRepository) GetAllHistory(id int) ([]_entities.Payment, error) {
	var history []_entities.Payment
	fmt.Println("historyRepository", history)

	tx := pr.database.Preload("User").Preload("Venue").Where("user_id = ?", id).Find(&history)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return history, nil
}

func (pr *PaymentRepository) GetByIdTransaction(id int) (_entities.Payment, error) {
	var transaction _entities.Payment

	err := pr.database.Where("id = ?", id).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (pr *PaymentRepository) CreateTransaction(booking _entities.Payment) (_entities.Payment, error) {
	err := pr.database.Save(&booking).Error

	if err != nil {
		return booking, err
	}

	return booking, nil
}

func (pr *PaymentRepository) UpdateTransaction(transaction _entities.Payment) (_entities.Payment, error) {
	err := pr.database.Save(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
