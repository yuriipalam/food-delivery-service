package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
	"time"
)

type CustomerRepositoryI interface {
	GetCustomerByID(int) (*model.Customer, error)
	GetCustomerByEmail(string) (*model.Customer, error)
	GetCustomerByPhone(string) (*model.Customer, error)
	CreateCustomer(*model.Customer) error
	DeleteCustomerByID(int) error
}

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (cr *CustomerRepository) GetCustomerByID(id int) (*model.Customer, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM customer WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("could not prepare a statement for id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot run query for id %d", id)
	}

	var customer model.Customer

	err = row.Scan(
		&customer.ID,
		&customer.Email,
		&customer.Password,
		&customer.Phone,
		&customer.FirstName,
		&customer.LastName,
		&customer.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan customer with id %d", id)
	}

	return &customer, nil
}

func (cr *CustomerRepository) GetCustomerByEmail(email string) (*model.Customer, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM customer WHERE email = $1")
	if err != nil {
		return nil, fmt.Errorf("could not prepare a statement for email %s", email)
	}

	row := stmt.QueryRow(email)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot run query for email %s", email)
	}

	var customer model.Customer

	err = row.Scan(
		&customer.ID,
		&customer.Email,
		&customer.Password,
		&customer.Phone,
		&customer.FirstName,
		&customer.LastName,
		&customer.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan customer with email %s", email)
	}

	return &customer, nil
}

func (cr *CustomerRepository) GetCustomerByPhone(phone string) (*model.Customer, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM customer WHERE phone = $1")
	if err != nil {
		return nil, fmt.Errorf("could not prepare a statement for phone %s", phone)
	}

	row := stmt.QueryRow(phone)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot run query for phone %s", phone)
	}

	var customer model.Customer

	err = row.Scan(
		&customer.ID,
		&customer.Email,
		&customer.Password,
		&customer.Phone,
		&customer.FirstName,
		&customer.LastName,
		&customer.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan customer with phone %s", phone)
	}

	return &customer, nil
}

func (cr *CustomerRepository) CreateCustomer(customer *model.Customer) error {
	stmt, err := cr.db.Prepare(`INSERT INTO customer (email, password, phone, first_name, last_name, created_at)
									  VALUES ($1, $2, $3, $4, $5, $6)
									  RETURNING id`)
	if err != nil {
		return fmt.Errorf("cannot prepare statement for the given customer model")
	}

	row := stmt.QueryRow(
		customer.Email,
		customer.Password,
		customer.Phone,
		customer.FirstName,
		customer.LastName,
		time.Now(),
	)
	if row.Err() != nil {
		return fmt.Errorf("cannot execute query for the given customer")
	}

	var lastInsertedID int
	err = row.Scan(&lastInsertedID)
	if err != nil {
		return fmt.Errorf("cannot scan last inserted id")
	}

	customer.ID = lastInsertedID

	return nil
}

func (cr *CustomerRepository) DeleteCustomerByID(id int) error {
	stmt, err := cr.db.Prepare("DELETE FROM customer WHERE id = $1")
	if err != nil {
		return fmt.Errorf("cannot prepare statement for id %d", id)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("cannot exec query for id %d", id)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affacted")
	}

	return nil
}
