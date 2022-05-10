package payment

import _entities "capstone/entities"

type PaymentRepositryInterface interface {
	GetAllHistory(id int) ([]_entities.Payment, error)
	GetOwnerHistory(id int) ([]_entities.Venue, error)
	CreateTransaction(booking _entities.Payment) (_entities.Payment, error)
	UpdateTransaction(transaction _entities.Payment) (_entities.Payment, error)
	GetByIdTransaction(ID int) (_entities.Payment, error)
}
