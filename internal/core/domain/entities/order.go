package entities

import (
	"log"
	valueobjects "tech-challenge-fase-1/internal/core/domain/value_objects"
)

type OrderStatus string

func (s OrderStatus) String() string {
	return string(s)
}

const (
	Pending          OrderStatus = "PENDING"
	Paid             OrderStatus = "PAID"
	Rejected         OrderStatus = "REJECTED"
	AWAITING_PAYMENT OrderStatus = "AWAITING_PAYMENT"
)

type Order struct {
	ID     string
	Items  valueobjects.OrderItems
	Status OrderStatus
}

func NewOrder(id string, items valueobjects.OrderItems, status OrderStatus) *Order {

	return &Order{ID: id, Items: items, Status: status}
}

func (o *Order) GetID() string {
	return o.ID
}

func (o *Order) GetItems() valueobjects.OrderItems {
	return o.Items
}

func (o *Order) GetStatus() OrderStatus {
	return o.Status
}

func (o *Order) GetTotal() *int64 {
	var total int64

	for _, item := range o.Items {
		total = total + *item.GetTotal()
	}

	return &total
}

func (o *Order) AddItem(product *Product, quantity *int) {
	amount := int64(product.GetPrice() * 100)
	productName := product.GetName()

	equal := false
	for index, item := range o.Items {

		log.Println(*item.GetProductName() == product.Name)
		if *item.GetProductName() == product.Name {

			quantity := *item.Quantity + *quantity
			item.Quantity = &quantity

			o.Items[index] = item
			equal = true
			break
		}

	}

	if !equal {
		o.Items = append(o.Items, valueobjects.OrderItem{
			Amount:      &amount,
			ProductName: &productName,
			Quantity:    quantity,
		})
	}

}
