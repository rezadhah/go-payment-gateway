package main

import (
	"go-payment-gateway/internal/rest"

	"github.com/labstack/echo/v4"
)

func main() {
	paymentHandler := rest.PaymentHandler()

	e := echo.New()
	e.POST("/payments", paymentHandler.Pay)
	e.Logger.Fatal(e.Start(":8080"))
}
