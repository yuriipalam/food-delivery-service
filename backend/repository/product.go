package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type ProductRepositoryI interface {
	GetProductByID(int) (*model.Product, error)
	GetAllProducts() ([]model.Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	stmt, err := pr.db.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("could not prepare statement for a product with id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot run a query for a given id %d", id)
	}

	var product model.Product

	err = row.Scan(
		&product.ID,
		&product.SupplierID,
		&product.CategoryID,
		&product.Name,
		&product.Image,
		&product.Description,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan product")
	}

	return &product, nil
}

func (pr *ProductRepository) GetAllProducts() ([]model.Product, error) {
	stmt, err := pr.db.Prepare("SELECT * FROM product")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("cannot execute query")
	}

	var products []model.Product

	for rows.Next() {
		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.SupplierID,
			&product.CategoryID,
			&product.Name,
			&product.Image,
			&product.Description,
			&product.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan product")
		}

		products = append(products, product)
	}

	return products, nil
}
