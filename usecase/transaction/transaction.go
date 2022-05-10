package payment

import (
	_entities "capstone/entities"
	_paymentRepo "capstone/repository/transaction"
	"capstone/usecase/payment"
	"fmt"
	"strconv"

	"github.com/jinzhu/copier"
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

func (pus *PaymentUseCase) GetOwnerHistory(id int) ([]_entities.VenueResponse, error) {
	venueResponse := []_entities.VenueResponse{}

	historyOwner, err := pus.paymentRepository.GetOwnerHistory(id)

	if err != nil {
		return venueResponse, err
	}

	copier.Copy(&venueResponse, &historyOwner)

	return venueResponse, nil
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

	fmt.Println("transaction", transaction)

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
		transaction.TotalPrice = transaction.TotalPrice + updatedTransaction.TotalPrice
	}

	fmt.Println("updatedTransaction", updatedTransaction)

	return nil
}
