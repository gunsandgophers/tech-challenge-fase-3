package valueobjects

type OrderItems []OrderItem

type OrderItem struct {
	Amount      *int64  `json:"amount,omitempty"`
	Quantity    *int    `json:"quantity,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
}

func (i *OrderItem) GetAmount() *int64 {
	return i.Amount
}

func (i *OrderItem) GetQuantity() *int {
	return i.Quantity
}

func (i *OrderItem) GetProductName() *string {
	return i.ProductName
}

func (i *OrderItem) GetTotal() *int64 {

	total := *i.Amount * int64(*i.Quantity)
	return &total
}
