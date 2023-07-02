package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
	"food_delivery/request"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CustomerRepositoryI interface {
	GetCustomerByID(int) (*model.Customer, error)
	GetCustomerByEmail(string) (*model.Customer, error)
	GetCustomerByPhone(string) (*model.Customer, error)
	CreateCustomer(registerRequest *request.RegisterRequest) (*model.Customer, error)
	UpdateCustomerByID(int, *request.UpdateCustomer) error
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
			return nil, fmt.Errorf("customer with id %d does not exist", id)
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
	if err := cr.checkIfEmailOrPhoneAlreadyExist(req.Email, req.Phone); err != nil {
		return nil, err
	}

	stmt, err := cr.db.Prepare(`INSERT INTO customer (email, password, phone, first_name, last_name, created_at)
									  VALUES ($1, $2, $3, $4, $5, $6)
									  RETURNING id`)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for the given customer model")
	}

	p, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

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

func (cr *CustomerRepository) UpdateCustomerByID(id int, customer *request.UpdateCustomer) error {
	customerFromDB, err := cr.checkIfCustomerExistByID(id)
	if err != nil {
		return err
	}

	anyChanges := false

	if customer.Phone != customerFromDB.Phone {
		if err := cr.updateField(id, "phone", customer.Phone); err != nil {
			return err
		}
		anyChanges = true
	}
	if customer.FirstName != customerFromDB.FirstName {
		if err := cr.updateField(id, "first_name", customer.FirstName); err != nil {
			return err
		}
		anyChanges = true
	}
	if customer.LastName != customerFromDB.LastName {
		if err := cr.updateField(id, "last_name", customer.LastName); err != nil {
			return err
		}
		anyChanges = true
	}


	//customerType := reflect.TypeOf(*customer)
	//customerValueOf := reflect.ValueOf(*customer)
	//customerFromDBValueOf := reflect.ValueOf(*customerFromDB)
	//
	//for i := 0; i < customerType.NumField(); i++ {
	//	customerVal := customerValueOf.Field(i).Interface()
	//	customerFromDBVal := customerFromDBValueOf.Field(i).Interface()
	//
	//	if !reflect.DeepEqual(customerVal, customerFromDBVal) && !utils.IsDefaultValue(customerVal) {
	//		fieldName := customerType.Field(i).Tag.Get("json")
	//
	//		if err := cr.updateField(id, fieldName, customerVal); err != nil {
	//			return err
	//		}
	//
	//		count++
	//	}
	//}

	if !anyChanges {
		return fmt.Errorf("all the fields in provided structs are the same as in db")
	}

	return nil
}

func (cr *CustomerRepository) DeleteCustomerByID(id int) error {
	if _, err := cr.checkIfCustomerExistByID(id); err != nil {
		return err
	}

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

func (cr *CustomerRepository) checkIfEmailOrPhoneAlreadyExist(email string, phone string) error {
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

func (cr *CustomerRepository) checkIfCustomerExistByID(id int) (*model.Customer, error) {
	customerFromDB, err := cr.GetCustomerByID(id)
	if err != nil {
		return nil, err
	} else if customerFromDB == nil {
		return nil, fmt.Errorf("customer with id %d not found", id)
	}

	return customerFromDB, nil
}
