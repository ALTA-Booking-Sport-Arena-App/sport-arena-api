package payment

import _entities "capstone/entities"

type PaymentUseCaseInterface interface {
	GetAllHistory(id int) ([]_entities.Payment, error)
	CreateBooking(booking _entities.Payment) (_entities.Payment, error)
	// CreatePayment(payment _entities.Payment) (_entities.Payment, error)
}
