// payment/service.go

package payment

import (
	"errors"
	"fmt"
	"go-payment-gateway/internal/method"
)

// PaymentService defines the payment-related methods.
type PaymentService interface {
	CreatePayment(request CreatePaymentRequest) (PaymentResponse, error)
}

type paymentService struct {
	paymentMethods func(name string) (method.PaymentProcessor, error)
}

// NewPaymentService creates a new PaymentService instance.
func NewPaymentService() PaymentService {
	paymentMethods := method.PaymentProcessFactory()
	return &paymentService{
		paymentMethods: paymentMethods,
	}
}

func (s *paymentService) CreatePayment(request CreatePaymentRequest) (PaymentResponse, error) {
	// Validate request (e.g., check amount, currency, etc.).
	if request.Amount <= 0 {
		return PaymentResponse{}, errors.New("invalid payment amount")
	}

	// Get the appropriate payment strategy based on the method.
	paymentMethod, err := s.paymentMethods(request.PaymentMethod)
	if err != nil {
		return PaymentResponse{}, err
	}

	paymentProcessorRequest := method.ProcessPaymentRequest{}

	// Process payment using the selected strategy.
	_, err = paymentMethod.ProcessPayment(paymentProcessorRequest)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf(ErrFailedProcessPayment, request.PaymentMethod, err)
	}

	return PaymentResponse{}, nil
}
