package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
	"food_delivery/request"
	"food_delivery/response"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type CustomerRepositoryI interface {
	GetCustomerByID(int) (*model.Customer, error)
	GetCustomerByEmail(string) (*model.Customer, error)
	GetCustomerByPhone(string) (*model.Customer, error)
	CreateCustomer(registerRequest *request.RegisterRequest) (*model.Customer, error)
	UpdateCustomerFirstNameByID(int, *request.UpdateCustomer, *model.Customer) error
	UpdateCustomerLastNameByID(int, *request.UpdateCustomer, *model.Customer) error
	UpdateCustomerPhoneByID(int, *request.UpdateCustomer, *model.Customer) error
	UpdateCustomerPasswordByID(int, *request.UpdateCustomerPassword, *model.Customer) error
	DeleteCustomerByID(int) error
	CheckIfEmailOrPhoneAlreadyExist(string, string) error
	TryToGetCustomerByID(int, http.ResponseWriter) (*model.Customer, bool)
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

func (cr *CustomerRepository) CreateCustomer(req *request.RegisterRequest) (*model.Customer, error) {
	stmt, err := cr.db.Prepare(`INSERT INTO customer (email, password, phone, first_name, last_name, created_at)
									  VALUES ($1, $2, $3, $4, $5, $6)
									  RETURNING id`)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for the given customer model")
	}

	p, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot generate bcrypt hash from password")
	}

	row := stmt.QueryRow(
		req.Email,
		p,
		req.Phone,
		req.FirstName,
		req.LastName,
		time.Now(),
	)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot execute query for the given customer")
	}

	var lastInsertedID int
	err = row.Scan(&lastInsertedID)
	if err != nil {
		return nil, fmt.Errorf("cannot scan last inserted id")
	}

	return cr.GetCustomerByID(lastInsertedID)
}

func (cr *CustomerRepository) UpdateCustomerFirstNameByID(id int, req *request.UpdateCustomer, customer *model.Customer) error {
	if err := cr.updateField(id, "first_name", req.FirstName); err != nil {
		return err
	}
	customer.FirstName = req.FirstName

	return nil
}

func (cr *CustomerRepository) UpdateCustomerLastNameByID(id int, req *request.UpdateCustomer, customer *model.Customer) error {
	if err := cr.updateField(id, "last_name", req.LastName); err != nil {
		return err
	}
	customer.LastName = req.LastName

	return nil
}

func (cr *CustomerRepository) UpdateCustomerPhoneByID(id int, req *request.UpdateCustomer, customer *model.Customer) error {
	if err := cr.updateField(id, "phone", req.Phone); err != nil {
		return err
	}
	customer.Phone = req.Phone

	return nil
}

func (cr *CustomerRepository) UpdateCustomerPasswordByID(id int, req *request.UpdateCustomerPassword, customer *model.Customer) error {
	p, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot generate bcrypt hash from password")
	}

	if err := cr.updateField(id, "password", p); err != nil {
		return err
	}

	customer.Password = string(p)

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

func (cr *CustomerRepository) CheckIfEmailOrPhoneAlreadyExist(email string, phone string) error {
	c, err := cr.GetCustomerByEmail(email)
	if err != nil {
		return err
	} else if c != nil {
		return fmt.Errorf("email %s already exist", email)
	}

	c, err = cr.GetCustomerByPhone(phone)
	if err != nil {
		return err
	} else if c != nil {
		return fmt.Errorf("phone %s already exist", phone)
	}

	return nil
}

func (cr *CustomerRepository) TryToGetCustomerByID(id int, w http.ResponseWriter) (*model.Customer, bool) {
	customer, err := cr.GetCustomerByID(id)
	if err != nil {
		response.SendInternalServerError(w, err)
		return nil, false
	} else if customer == nil {
		response.SendNotFoundError(w, fmt.Errorf("customer with id %d not found", id))
		return nil, false
	}

	return customer, true
}

func (cr *CustomerRepository) updateField(id int, fieldName string, fieldVal any) error {
	stmtStr := fmt.Sprintf("UPDATE customer SET %s = $1 WHERE id = $2", fieldName)
	stmt, err := cr.db.Prepare(stmtStr)
	if err != nil {
		return fmt.Errorf("cannot prepare statement for id %d, fieldname %s", id, fieldName)
	}

	_, err = stmt.Exec(fieldVal, id)
	if err != nil {
		return fmt.Errorf("cannot execute query for id %d, fieldname %s", id, fieldName)
	}

	return nil
}
