package method

import "go-payment-gateway/payment"

type creditCard struct {
}

func CreditCardStrategy() payment.PaymentStrategy {
	return &creditCard{}
}

func (cc *creditCard) Pay(request payment.PaymentRequest) (payment.PaymentResponse, error) {
	return payment.PaymentResponse{}, nil
}
