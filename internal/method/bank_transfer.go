package method

import "go-payment-gateway/payment"

type bankTransfer struct {
}

func BankTransferStrategy() payment.PaymentStrategy {
	return &bankTransfer{}
}

func (bt *bankTransfer) Pay(request payment.PaymentRequest) (payment.PaymentResponse, error) {
	return payment.PaymentResponse{}, nil
}
