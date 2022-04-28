package payment

import (
	"capstone/delivery/helper"
	_middlewares "capstone/delivery/middlewares"
	_entities "capstone/entities"
	_payment "capstone/usecase/payment"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentUseCase _payment.PaymentUseCaseInterface
}

func NewFacilityHandler(paymentUseCase _payment.PaymentUseCaseInterface) PaymentHandler {
	return PaymentHandler{paymentUseCase}
}

func (ph *PaymentHandler) GetAllHistoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}

		history, err := ph.paymentUseCase.GetAllHistory(idToken)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get all histories"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all histories", history))
	}
}

func (ph *PaymentHandler) CreateFacilityHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var booking _entities.Payment

		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}

		booking.UserID = idToken

		fmt.Println("booking-handler", booking)

		err := c.Bind(&booking)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		_, err = ph.paymentUseCase.CreateBooking(booking)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("created booking failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("created booking successfully"))
	}
}
