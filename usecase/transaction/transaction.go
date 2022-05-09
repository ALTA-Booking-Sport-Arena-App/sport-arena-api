package payment

import (
	_entities "capstone/entities"
	_paymentRepo "capstone/repository/transaction"
	"capstone/usecase/payment"
	"strconv"
)

type PaymentUseCase struct {
	paymentRepository _paymentRepo.PaymentRepositryInterface
	paymentService    payment.Service
}

func NewPaymentUseCase(paymentRepo _paymentRepo.PaymentRepositryInterface, paymentService payment.Service) PaymentUseCaseInterface {
	return &PaymentUseCase{
		paymentRepository: paymentRepo,
		paymentService:    paymentService,
	}
}

func (pus *PaymentUseCase) GetAllHistory(id int) ([]_entities.Payment, error) {
	history, err := pus.paymentRepository.GetAllHistory(id)

	return history, err
}

func (pus *PaymentUseCase) CreateTransaction(booking _entities.Payment) (_entities.Payment, error) {
	booking.Status = "pending"

	newBooking, err := pus.paymentRepository.CreateTransaction(booking)

	if err != nil {
		return newBooking, err
	}

	paymentURL, err := pus.paymentService.GetPaymentURL(newBooking, booking.User)

	if err != nil {
		return newBooking, err
	}

	newBooking.PaymentURL = paymentURL

	newBooking, err = pus.paymentRepository.UpdateTransaction(newBooking)

	if err != nil {
		return newBooking, err
	}

	return newBooking, nil
}

func (pus *PaymentUseCase) ProcessPayment(input _entities.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := pus.paymentRepository.GetByIdTransaction(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := pus.paymentRepository.UpdateTransaction(transaction)

	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		updatedTransaction.TotalPrice += updatedTransaction.TotalPrice
	}

	return nil
}
