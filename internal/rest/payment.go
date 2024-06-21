package rest

import (
	"go-payment-gateway/internal/payment"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentService payment.PaymentService
}

func PaymentHandler() paymentHandler {
	svc := payment.NewPaymentService()
	return paymentHandler{paymentService: svc}
}

func (ph paymentHandler) Pay() func(c echo.Context) error {
	req := payment.PaymentRequest{}
	_, err := ph.paymentService.ProcessPayment(req)
	if err != nil {
		return nil
	}

	// return c.JSON(http.StatusOK, resp)
	return nil
}
