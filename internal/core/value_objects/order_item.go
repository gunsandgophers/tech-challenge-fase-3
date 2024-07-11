package valueobjects

type OrderItem struct {
	amount      int64
	quantity    int
	productName string
}

func NewOrderItem(amount int64, quantity int, productName string) *OrderItem {
	return &OrderItem{
		amount: amount,
		quantity: quantity,
		productName: productName,
	}
}

func (i *OrderItem) GetAmount() int64 {
	return i.amount
}

func (i *OrderItem) GetQuantity() int {
	return i.quantity
}

func (i *OrderItem) SetQuatity(quantity int) {
	i.quantity = quantity
}

func (i *OrderItem) GetProductName() string {
	return i.productName
}

func (i *OrderItem) GetTotal() int64 {
	return i.amount * int64(i.quantity)
}
