package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type SupplierRepositoryI interface {
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

	if rows.Err() != nil {
		return nil, fmt.Errorf("type mismatch on scanning rows")
	}

	return suppliers, nil
}
