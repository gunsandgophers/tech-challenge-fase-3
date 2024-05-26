package dtos

type MethodType string

const (
	PIX         MethodType = "PIX"
	CREDIT_CARD MethodType = "CREDIT_CARD"
)

type CheckoutDTO struct {
	PaymentLink *string
	Method      *MethodType
	Amount      *float64
}
