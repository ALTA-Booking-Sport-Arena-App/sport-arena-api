package payment

import _entities "capstone/entities"

type PaymentUseCaseInterface interface {
	GetAllHistory(id int) ([]_entities.Payment, error)
}
