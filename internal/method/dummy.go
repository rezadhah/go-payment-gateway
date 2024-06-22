package method

type dummy struct {
}

func DummyStrategy() *dummy {
	return &dummy{}
}

func (d *dummy) ProcessPayment(request ProcessPaymentRequest) (ProcessPaymentResponse, error) {
	return ProcessPaymentResponse{}, nil
}
