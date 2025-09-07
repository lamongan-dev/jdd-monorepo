package rest

import (
	"api-ticketing/domain"

	"github.com/labstack/echo/v4"
)

type IPaymuService interface {
	ProcessPayment(payment domain.PaymentBody) (*domain.PaymentResponse, error)
}

type PaymentHandler struct {
	Service IPaymuService
}

func (h *PaymentHandler) ProcessPayment(c echo.Context) error {
	var paymentBody domain.PaymentBody
	if err := c.Bind(&paymentBody); err != nil {
		return err
	}

	response, err := h.Service.ProcessPayment(paymentBody)
	if err != nil {
		return err
	}

	return c.JSON(201, response)
}

func NewPaymentHandler(e *echo.Group, svc IPaymuService) {
	handler := &PaymentHandler{
		Service: svc,
	}

	paymentGroup := e.Group("/payment")
	paymentGroup.POST("", handler.ProcessPayment)
}
