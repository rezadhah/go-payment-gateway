package method

import (
	"fmt"
)

const (
	BankTransfer = "bank_transfer"
	CreditCard   = "credit_card"
)

type PaymentProcessor interface {
	ProcessPayment(request ProcessPaymentRequest) (ProcessPaymentResponse, error)
}

func PaymentProcessFactory() func(name string) (PaymentProcessor, error) {
	return func(name string) (PaymentProcessor, error) {
		switch name {
		case BankTransfer:
			return BankTransferStrategy(), nil
		case CreditCard:
			return CreditCardStrategy(), nil
		default:
			return nil, fmt.Errorf(ErrUnknownPaymentMethod, name)
		}
	}
}
