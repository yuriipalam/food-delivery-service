package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type OrderRepositoryI interface {
	GetOrdersByCustomerID(int) ([]model.Order, error)
	GetSupplierNamesByOrderID(int) ([]int, []string, error)
	GetCustomerNameByOrderID(int) (string, error)
	GetProductNamesByOrderID(int) ([]int, []string, error)
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

func (or *OrderRepository) GetSupplierNamesByOrderID(id int) ([]int, []string, error) {
	query := `SELECT s.id, s.name FROM supplier s 
    		  JOIN order_supplier os ON s.id = os.supplier_id
    		  JOIN "order" o ON o.id = os.order_id
    		  WHERE o.id = $1`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot prepare statement for order_id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot run query for order_id %d", id)
	}

	var supplierIDs []int
	var supplierNames []string

	for rows.Next() {
		var supplierID int
		var supplierName string

		err := rows.Scan(
			&supplierID,
			&supplierName,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot scan supplier for order_id %d", id)
		}

		supplierIDs = append(supplierIDs, supplierID)
		supplierNames = append(supplierNames, supplierName)
	}

	return supplierIDs, supplierNames, nil
}

func (or *OrderRepository) GetCustomerNameByOrderID(id int) (string, error) {
	query := `SELECT (first_name || ' '|| last_name) AS customer_name FROM customer WHERE id = (SELECT customer_id FROM "order" WHERE id = $1)`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("cannot prepare query for order_id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return "", fmt.Errorf("cannot run query for order_id %d", id)
	}

	var name string

	err = row.Scan(&name)
	if err != nil {
		return "", fmt.Errorf("cannot scan for order_id %d", id)
	}

	return name, nil
}

func (or *OrderRepository) GetProductNamesByOrderID(id int) ([]int, []string, error) {
	query := `SELECT p.id, p.name FROM product p 
    		  JOIN order_product ps ON p.id = ps.product_id
    		  JOIN "order" o ON o.id = ps.order_id
    		  WHERE o.id = $1`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot prepare statement for order_id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot run query for order_id %d", id)
	}

	var productIDs []int
	var productNames []string

	for rows.Next() {
		var productID int
		var productName string

		err := rows.Scan(
			&productID,
			&productName,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot scan product for order_id %d", id)
		}

		productIDs = append(productIDs, productID)
		productNames = append(productNames, productName)
	}

	return productIDs, productNames, nil
}
