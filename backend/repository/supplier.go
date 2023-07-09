package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type SupplierRepositoryI interface {
	GetSupplierByID(int) (*model.Supplier, error)
	GetSuppliersByCategoryIDs([]int) ([]model.Supplier, error)
	GetAllSuppliers() ([]model.Supplier, error)
	GetCategoryNamesBySupplierID(int) ([]int, []string, error)
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
		return nil, fmt.Errorf("cannot execute query for %d id", id)
	}

	var supplier model.Supplier

	err = row.Scan(
		&supplier.ID,
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
		return nil, fmt.Errorf("cannot scan supplier")
	}

	return &supplier, nil
}

func (sr *SupplierRepository) GetSuppliersByCategoryIDs(ids []int) ([]model.Supplier, error) {
	// many-to-many relation, shortly, this query equivalent to SELECT * FROM supplier WHERE category_id = $1
	query := `SELECT *
			  FROM supplier s
			  WHERE EXISTS(SELECT *
			  			 FROM supplier_category sc
			  			 WHERE ARRAY[$1::integer[]] <@ (SELECT ARRAY_AGG(DISTINCT category_id)
			  									FROM supplier_category
			  									WHERE supplier_id = sc.supplier_id)
			  			   AND sc.supplier_id = s.id);`

	stmt, err := sr.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for %v category_ids", ids)
	}

	rows, err := stmt.Query(pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("cannot execute query for %v category_ids", ids)
	}

	var suppliers []model.Supplier

	for rows.Next() {
		var supplier model.Supplier

		err = rows.Scan(
			&supplier.ID,
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

func (sr *SupplierRepository) GetCategoryNamesBySupplierID(id int) ([]int,[]string, error) {
	stmt, err := sr.db.Prepare("SELECT category_id FROM supplier_category WHERE supplier_id = $1")
	if err != nil {
		return nil, nil, fmt.Errorf("cannot prepare statement for supplier id %d", id)
	}

	var categoryIDs []int

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot run query for supplier id %d", id)
	}

	for rows.Next() {
		var categoryID int

		if err := rows.Scan(&categoryID); err != nil {
			return nil, nil, fmt.Errorf("cannot scan category id")
		}

		categoryIDs = append(categoryIDs, categoryID)
	}

	stmt, err = sr.db.Prepare("SELECT name FROM category WHERE id = ANY($1)")
	if err != nil {
		return nil, nil, fmt.Errorf("cannot prepare statement for categories %v", categoryIDs)
	}

	rows, err = stmt.Query(pq.Array(categoryIDs))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot run query for category id %d", id)
	}

	var names []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, nil, fmt.Errorf("cannot scan name")
		}

		names = append(names, name)
	}

	return categoryIDs, names, nil
}
