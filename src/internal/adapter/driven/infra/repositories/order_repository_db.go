package repositories

import (
	"encoding/json"
	"tech-challenge-fase-1/internal/adapter/driven/infra/database"
	"tech-challenge-fase-1/internal/core/domain/entities"
	valueobjects "tech-challenge-fase-1/internal/core/domain/value_objects"

	"github.com/google/uuid"
)

type OrderRepositoryDB struct {
	conn database.ConnectionDB
}

func NewOrderRepositoryDB(conn database.ConnectionDB) *OrderRepositoryDB {
	return &OrderRepositoryDB{conn: conn}
}

func (r *OrderRepositoryDB) Open(customerID *string) (*string, error) {

	var id = uuid.NewString()
	sql := "INSERT INTO orders(id, customer_id) VALUES($1, $2)"

	err := r.conn.Exec(sql, id, customerID)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *OrderRepositoryDB) FindByID(orderID *string) (*entities.Order, error) {

	sql := `SELECT id, to_jsonb(items), status FROM orders WHERE id = $1`

	row := r.conn.QueryRow(sql, orderID)

	return r.toEntity(row)
}

func (r *OrderRepositoryDB) AddOrderItem(order *entities.Order) error {

	sql := `UPDATE orders SET items = $1  WHERE id = $2;`

	err := r.conn.Exec(sql, order.GetItems(), order.ID)

	return err
}

func (r *OrderRepositoryDB) UpdateStatus(order *entities.Order) error {

	sql := `UPDATE orders SET status = $1  WHERE id = $2;`

	err := r.conn.Exec(sql, order.GetStatus(), order.ID)

	return err
}

func (r *OrderRepositoryDB) toEntity(row database.RowDB) (*entities.Order, error) {
	var id string
	var item valueobjects.OrderItems
	var itemByte []byte
	var status entities.OrderStatus

	err := row.Scan(&id, &itemByte, &status)
	if err != nil {
		if err.Error() == ErrNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	err = json.Unmarshal(itemByte, &item)
	if err != nil {
		return nil, err
	}

	return entities.NewOrder(id, item, status), nil
}
