package payment

import (
	_entities "capstone/entities"
	_paymentRepo "capstone/repository/payment"
	"fmt"

	"github.com/jinzhu/copier"
)

type PaymentUseCase struct {
	paymentRepository _paymentRepo.PaymentRepositryInterface
}

func NewFacilityUseCase(paymentRepo _paymentRepo.PaymentRepositryInterface) PaymentUseCaseInterface {
	return &PaymentUseCase{
		paymentRepository: paymentRepo,
	}
}

func (pus *PaymentUseCase) GetAllHistory(id int) ([]_entities.Payment, error) {
	var historyResponse []_entities.PaymentResponse

	history, err := pus.paymentRepository.GetAllHistory(id)

	fmt.Println("historyUseCase", history)

	copier.Copy(&historyResponse, &history)

	fmt.Println("historyResponseUseCase", historyResponse)
	return history, err
}
