package payment

import (
	"fmt"
)

type PaymentStrategy interface {
	Pay(request PaymentRequest) (PaymentResponse, error)
}

func StrategyFactory() func(name string) (PaymentStrategy, error) {
	return func(name string) (PaymentStrategy, error) {
		switch name {
		default:
			return nil, fmt.Errorf(ErrUnknownPaymentMethod, name)
		}
	}
}
