package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type SupplierRepositoryI interface {
	GetSupplierByID(int) (*model.Supplier, error)
	GetSuppliersByCategoryID(int) ([]model.Supplier, error)
	GetAllSuppliers() ([]model.Supplier, error)
}

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (sr *SupplierRepository) GetSupplierByID(id int) (*model.Supplier, error) {
	stmt, err := sr.db.Prepare("SELECT * FROM supplier WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for %d id", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot execute queryRow for %d id", id)
	}

	var supplier model.Supplier

	err = row.Scan(
		&supplier.ID,
		&supplier.CategoryID,
		&supplier.Name,
		&supplier.Image,
		&supplier.Description,
		&supplier.TimeOpening,
		&supplier.TimeClosing,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan product")
	}

	return &supplier, nil
}

func (sr *SupplierRepository) GetSuppliersByCategoryID(id int) ([]model.Supplier, error) {
	stmt, err := sr.db.Prepare("SELECT * FROM supplier WHERE category_id = $1")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for %d category_id", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query for category_id %d", id)
	}

	var suppliers []model.Supplier

	for rows.Next() {
		var supplier model.Supplier

		err = rows.Scan(
			&supplier.ID,
			&supplier.CategoryID,
			&supplier.Name,
			&supplier.Image,
			&supplier.Description,
			&supplier.TimeOpening,
			&supplier.TimeClosing,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan product")
		}

		suppliers = append(suppliers, supplier)
	}

	if len(suppliers) == 0 {
		return nil, nil
	}

	return suppliers, nil
}

func (sr *SupplierRepository) GetAllSuppliers() ([]model.Supplier, error) {
	stmt, err := sr.db.Prepare("SELECT * FROM supplier")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("cannot execute query")
	}

	var suppliers []model.Supplier

	for rows.Next() {
		var supplier model.Supplier

		err := rows.Scan(
			&supplier.ID,
			&supplier.CategoryID,
			&supplier.Name,
			&supplier.Image,
			&supplier.Description,
			&supplier.TimeOpening,
			&supplier.TimeClosing,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan supplier")
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}
