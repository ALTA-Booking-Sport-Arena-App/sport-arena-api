package payment

import (
	"capstone/delivery/helper"
	_middlewares "capstone/delivery/middlewares"
	_payment "capstone/usecase/payment"
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
