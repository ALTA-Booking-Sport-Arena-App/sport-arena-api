package payment

import (
	"capstone/delivery/helper"
	_middlewares "capstone/delivery/middlewares"
	_entities "capstone/entities"
	_payment "capstone/usecase/transaction"
	_userUseCase "capstone/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentUseCase _payment.PaymentUseCaseInterface
	userUseCase    _userUseCase.UserUseCaseInterface
}

func NewPaymentHandler(paymentUseCase _payment.PaymentUseCaseInterface, userUseCase _userUseCase.UserUseCaseInterface) PaymentHandler {
	return PaymentHandler{paymentUseCase, userUseCase}
}

func (ph *PaymentHandler) GetAllHistoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		history, err := ph.paymentUseCase.GetAllHistory(idToken)

		historyResponse := []map[string]interface{}{}
		for i := 0; i < len(history); i++ {
			response := map[string]interface{}{
				"id": history[i].ID,
				"venue": map[string]interface{}{
					"name":       history[i].Venue.Name,
					"location":   history[i].Venue.Address,
					"price":      history[i].TotalPrice,
					"image":      history[i].Venue.Image,
					"status":     history[i].Status,
					"day":        history[i].Day,
					"start_date": history[i].StartDate,
					"end_date":   history[i].EndDate,
				},
			}
			historyResponse = append(historyResponse, response)
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed get all histories", http.StatusBadRequest))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all histories", http.StatusOK, historyResponse))
	}
}

func (ph *PaymentHandler) CreateBookingHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var booking _entities.Payment

		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		userProfile, _ := ph.userUseCase.GetUserProfile(idToken)

		booking.UserID = uint(idToken)
		booking.User = userProfile

		errBind := c.Bind(&booking)

		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errBind.Error(), http.StatusBadRequest))
		}
		newBooking, errCreate := ph.paymentUseCase.CreateTransaction(booking)

		response := map[string]interface{}{
			"payment_url": newBooking.PaymentURL,
		}

		if errCreate != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("created booking failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("created booking successfully", http.StatusOK, response))
	}
}

func (ph *PaymentHandler) GetNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input _entities.TransactionNotificationInput

		err := c.Bind(&input)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to bind data", http.StatusBadRequest))
		}

		err = ph.paymentUseCase.ProcessPayment(input)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to process notification", http.StatusBadRequest))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Failed to process notification", http.StatusOK))
	}
}
