package entities

import (
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"

	"github.com/google/uuid"
)

type OrderStatus string

func (s OrderStatus) String() string {
	return string(s)
}

const (
	PENDING          OrderStatus = "PENDING"
	PAID             OrderStatus = "PAID"
	REJECTED         OrderStatus = "REJECTED"
	AWAITING_PAYMENT OrderStatus = "AWAITING_PAYMENT"
)

type Order struct {
	id     string
	customerId string
	items  []*valueobjects.OrderItem
	status OrderStatus
}

func CreateOpenOrder(customerId string) *Order {
	return RestoreOrder(
		uuid.NewString(),
		customerId,
		make([]*valueobjects.OrderItem, 0),
		PENDING,
	)
}

func RestoreOrder(id string, customerId string, items []*valueobjects.OrderItem, status OrderStatus) *Order {
	return &Order{
		id: id,
		customerId: customerId,
		items: items,
		status: status,
	}
}

func (o *Order) GetId() string {
	return o.id
}

func (o *Order) GetCustomerId() string {
	return o.customerId
}

func (o *Order) GetItems() []*valueobjects.OrderItem {
	return o.items
}

func (o *Order) GetStatus() OrderStatus {
	return o.status
}

func (o *Order) SetStatus(status OrderStatus) {
	o.status = status
}

func (o *Order) GetTotal() int64 {
	var total int64
	for _, item := range o.items {
		total = total + item.GetTotal()
	}
	return total
}

func (o *Order) FindOrderItem(productName string) *valueobjects.OrderItem {
	for _, item := range o.items {
		if item.GetProductName() == productName {
			return item
		}
	}
	return nil
}

func (o *Order) AddItem(product *Product, quantity int) {
	amount := int64(product.GetPrice() * 100)
	productName := product.GetName()
	item := o.FindOrderItem(productName)
	if item == nil {
		item = valueobjects.NewOrderItem(amount, 0, productName)
		o.items = append(o.items, item)
	}
	quantity = item.GetQuantity() + quantity
	item.SetQuatity(quantity)
}
