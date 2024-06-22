package method

type bankTransfer struct {
}

func BankTransferStrategy() *bankTransfer {
	return &bankTransfer{}
}

func (bt *bankTransfer) ProcessPayment(request ProcessPaymentRequest) (ProcessPaymentResponse, error) {
	return ProcessPaymentResponse{}, nil
}
