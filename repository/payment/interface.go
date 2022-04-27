package payment

import _entities "capstone/entities"

type PaymentRepositryInterface interface {
	GetAllHistory(id int) ([]_entities.Payment, error)
}
