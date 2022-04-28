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

	tx := pr.database.Find(&history)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return history, nil
}

func (pr *PaymentRepository) CreateBooking(booking _entities.Payment) (_entities.Payment, error) {
	fmt.Println("booking-repo", booking)

	tx := pr.database.Save(&booking)

	if tx.Error != nil {
		return booking, tx.Error
	}

	return booking, nil
}
