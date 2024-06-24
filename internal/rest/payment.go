package rest

import (
	"go-payment-gateway/internal/payment"
	"net/http"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentService payment.PaymentService
}

func PaymentHandler() paymentHandler {
	svc := payment.NewPaymentService()
	return paymentHandler{paymentService: svc}
}

func (ph paymentHandler) Pay(c echo.Context) error {
	req := payment.CreatePaymentRequest{}
	resp, err := ph.paymentService.CreatePayment(req)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, resp)
}
