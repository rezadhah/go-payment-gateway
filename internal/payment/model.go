package payment

type (

	// CreatePaymentRequest represents a create payment request from the client.
	CreatePaymentRequest struct {
		Amount        float64
		Currency      string
		PaymentType   string // CreditCard, BankTransfer, etc.
		PaymentMethod string // Visa, Mastercard, etc.
	}

	// PaymentResponse represents the response after processing the payment.
	PaymentResponse struct {
		Success bool
		Message string
	}
)
