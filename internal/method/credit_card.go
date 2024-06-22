package method

type creditCard struct {
}

func CreditCardStrategy() *creditCard {
	return &creditCard{}
}

func (cc *creditCard) ProcessPayment(request ProcessPaymentRequest) (ProcessPaymentResponse, error) {
	return ProcessPaymentResponse{}, nil
}
