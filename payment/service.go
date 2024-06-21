// payment/service.go

package payment

import (
	"errors"
)

// PaymentService defines the payment-related methods.
type PaymentService interface {
	ProcessPayment(request PaymentRequest) (PaymentResponse, error)
}

type paymentService struct {
	paymentMethods func(name string) (PaymentStrategy, error)
}

// NewPaymentService creates a new PaymentService instance.
func NewPaymentService() PaymentService {
	paymentMethods := StrategyFactory()
	return &paymentService{
		paymentMethods: paymentMethods,
	}
}

func (s *paymentService) ProcessPayment(request PaymentRequest) (PaymentResponse, error) {
	// Validate request (e.g., check amount, currency, etc.).
	if request.Amount <= 0 {
		return PaymentResponse{}, errors.New("invalid payment amount")
	}

	// Get the appropriate payment strategy based on the method.
	paymentMethod, err := s.paymentMethods(request.PaymentMethod)
	if err != nil {
		return PaymentResponse{}, err
	}

	// Process payment using the selected strategy.
	return paymentMethod.Pay(request)
}
