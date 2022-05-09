package payment

import _entities "capstone/entities"

type PaymentUseCaseInterface interface {
	GetAllHistory(id int) ([]_entities.Payment, error)
	CreateTransaction(booking _entities.Payment) (_entities.Payment, error)
	ProcessPayment(input _entities.TransactionNotificationInput) error
}
