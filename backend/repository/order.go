package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type OrderRepositoryI interface {
	GetOrdersByCustomerID(int) ([]model.Order, error)
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) GetOrdersByCustomerID(id int) ([]model.Order, error) {
	stmt, err := or.db.Prepare("SELECT * FROM \"order\" WHERE customer_id = $1")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for customer_id id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("cannot run query for customer_id id %d", id)
	}

	var orders []model.Order

	for rows.Next() {
		var order model.Order

		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.RecipientFullName,
			&order.Address,
			&order.Price,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan order for customer_id id %d", id)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
