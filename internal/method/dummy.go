package method

import "go-payment-gateway/payment"

type dummy struct {
}

func DummyStrategy() payment.PaymentStrategy {
	return &dummy{}
}

func (d *dummy) Pay(request payment.PaymentRequest) (payment.PaymentResponse, error) {
	return payment.PaymentResponse{}, nil
}
