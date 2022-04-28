package payment

import _entities "capstone/entities"

type PaymentRepositryInterface interface {
	GetAllHistory(id int) ([]_entities.Payment, error)
	CreateBooking(booking _entities.Payment) (_entities.Payment, error)
	// CreatePayment(payment _entities.Payment) (_entities.Payment, error)
}
