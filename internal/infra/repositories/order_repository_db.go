package repositories

import (
	"encoding/json"
	"tech-challenge-fase-1/internal/infra/database"
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
)

type OrderRepositoryDB struct {
	conn database.ConnectionDB
}

func NewOrderRepositoryDB(conn database.ConnectionDB) *OrderRepositoryDB {
	return &OrderRepositoryDB{conn: conn}
}

func (r *OrderRepositoryDB) Insert(order *entities.Order) error {
	sql := `
	INSERT INTO orders(id, customer_id, items, status)
	VALUES ($1, $2, $3, $4)
	`
	var customerId *string
	if order.GetCustomerId() != "" {
		*customerId = order.GetCustomerId()
	}

	return r.conn.Exec(
		sql,
		order.GetId(),
		customerId,
		newOrderItemHelperList(order.GetItems()),
		order.GetStatus(),
	)
}

func (r *OrderRepositoryDB) FindByID(orderId *string) (*entities.Order, error) {
	sql := `SELECT id, to_jsonb(items), status FROM orders WHERE id = $1`
	row := r.conn.QueryRow(sql, orderId)
	return r.toEntity(row)
}

// func (r *OrderRepositoryDB) AddOrderItem(order *entities.Order) error {
// 	sql := `UPDATE orders SET items = $1  WHERE id = $2;`
// 	err := r.conn.Exec(sql, order.GetItems(), order.GetId())
// 	return err
// }
//
// func (r *OrderRepositoryDB) UpdateStatus(order *entities.Order) error {
// 	sql := `UPDATE orders SET status = $1  WHERE id = $2;`
// 	err := r.conn.Exec(sql, order.GetStatus(), order.GetId())
// 	return err
// }

func (r *OrderRepositoryDB) toEntity(row database.RowDB) (*entities.Order, error) {
	var id string
	var customerId *string
	var items []*valueobjects.OrderItem
	var itemByte []byte
	var status entities.OrderStatus
	err := row.Scan(&id, &customerId, &itemByte, &status)
	if err != nil {
		if err.Error() == ErrNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	err = json.Unmarshal(itemByte, &items)
	if err != nil {
		return nil, err
	}
	return entities.RestoreOrder(id, *customerId, items, status), nil
}

func orderItemsFromHelper(orderItemsHelper []*orderItemHelper) []*valueobjects.OrderItem {
	orderItems := make([]*valueobjects.OrderItem, 0)
	for _, item := range orderItemsHelper {
		orderItems = append(
			orderItems,
			valueobjects.NewOrderItem(item.Amount, item.Quantity, item.ProductName),
		)
	}
	return orderItems
}

func newOrderItemHelperList(orderItems []*valueobjects.OrderItem) []*orderItemHelper {
	orderItemsHelper := make([]*orderItemHelper, 0)
	for _, item := range orderItems {
		orderItemsHelper = append(
			orderItemsHelper,
			&orderItemHelper{
				Amount: item.GetAmount(),
				Quantity: item.GetQuantity(),
				ProductName: item.GetProductName(),
			},
		)
	}
	return orderItemsHelper
}

type orderItemHelper struct {
	Amount      int64  `json:"amount,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	ProductName string `json:"product_name,omitempty"`
}
