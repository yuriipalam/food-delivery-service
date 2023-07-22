package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
	"food_delivery/request"
	"food_delivery/response"
	_ "github.com/lib/pq"
	"time"
)

type OrderRepositoryI interface {
	GetOrderByID(int) (*model.Order, error)
	GetOrdersByCustomerID(int) ([]model.Order, error)
	CreateOrder(*request.CreateOrderRequest) (*model.Order, error)
	GetSupplierResponsesByOrderID(int) ([]response.OrderSupplierResponse, error)
	GetCustomerNameByOrderID(int) (string, error)
	GetProductResponsesByOrderID(int) ([]response.OrderProductResponse, error)
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) GetOrderByID(id int) (*model.Order, error) {
	stmt, err := or.db.Prepare("SELECT * FROM \"order\" WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for order_id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot run query for order_id %d", id)
	}

	var order model.Order

	err = row.Scan(
		&order.ID,
		&order.CustomerID,
		&order.RecipientFullName,
		&order.Address,
		&order.Price,
		&order.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot scan order with id %d", id)
	}

	return &order, nil
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

func (or *OrderRepository) CreateOrder(req *request.CreateOrderRequest) (*model.Order, error) {
	query := `INSERT INTO "order" (customer_id, recipient_full_name, address, price, created_at)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stmt, err := or.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare insert statement for order")
	}

	var lastInsertedID int

	row := stmt.QueryRow(
		req.CustomerID,
		req.RecipientFullName,
		req.Address,
		req.Price,
		time.Now(),
	)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot execute statement for order")
	}

	err = row.Scan(&lastInsertedID)
	if err != nil {
		return nil, fmt.Errorf("cannot get last inserted id for order")
	}

	for _, supplierID := range req.SupplierIDs {
		err := or.insertIntoOrderSupplier(lastInsertedID, supplierID)
		if err != nil {
			return nil, err
		}
	}

	for _, pair := range req.Products {
		err := or.insertIntoOrderProduct(lastInsertedID, pair.ProductID, pair.ProductQuantity)
		if err != nil {
			return nil, err
		}
	}

	order, err := or.GetOrderByID(lastInsertedID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (or *OrderRepository) GetSupplierResponsesByOrderID(id int) ([]response.OrderSupplierResponse, error) {
	query := `SELECT s.id, s.name, s.image FROM supplier s 
    		  JOIN order_supplier os ON s.id = os.supplier_id
    		  JOIN "order" o ON o.id = os.order_id
    		  WHERE o.id = $1`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for order_id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("cannot run query for order_id %d", id)
	}

	var suppliers []response.OrderSupplierResponse

	for rows.Next() {
		var supplier response.OrderSupplierResponse

		var imageName string
		err := rows.Scan(
			&supplier.SupplierID,
			&supplier.SupplierName,
			&imageName,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan supplier for order_id %d", id)
		}

		supplier.SupplierImageURL = fmt.Sprintf("http://localhost:8080/images/suppliers/%s", imageName)

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
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

func (or *OrderRepository) GetProductResponsesByOrderID(id int) ([]response.OrderProductResponse, error) {
	query := `SELECT p.id, p.name, p.supplier_id, os.product_quantity FROM product p 
    		  JOIN order_product os ON p.id = os.product_id
    		  JOIN "order" o ON o.id = os.order_id
    		  WHERE o.id = $1`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for order_id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("cannot run query for order_id %d", id)
	}

	var products []response.OrderProductResponse

	for rows.Next() {
		var product response.OrderProductResponse

		err := rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.ProductSupplierID,
			&product.ProductQuantity,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan product for order_id %d", id)
		}

		products = append(products, product)
	}

	return products, nil
}

func (or *OrderRepository) insertIntoOrderProduct(orderID int, productID int, productQuantity int) error {
	query := `INSERT INTO order_product (order_id, product_id, product_quantity)
			  VALUES ($1, $2, $3)`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("cannot prepare insert query for order_id %d", orderID)
	}

	row := stmt.QueryRow(orderID, productID, productQuantity)
	if row.Err() != nil {
		return fmt.Errorf("cannot run query for order_id %d", orderID)
	}

	return nil
}

func (or *OrderRepository) insertIntoOrderSupplier(orderID int, supplierID int) error {
	query := `INSERT INTO order_supplier (order_id, supplier_id)
			  VALUES ($1, $2)`

	stmt, err := or.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("cannot prepare insert query for order_id %d", orderID)
	}

	row := stmt.QueryRow(orderID, supplierID)
	if row.Err() != nil {
		return fmt.Errorf("cannot run query for order_id %d", orderID)
	}

	return nil
}
