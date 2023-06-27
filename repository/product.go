package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type ProductRepositoryI interface {
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

func (cr *ProductRepository) GetAllProducts() ([]model.Product, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM product")
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
			&product.Name,
			&product.Image,
			&product.Description,
			&product.Ingredients,
			&product.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan product")
		}

		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("type mismatch on scanning rows")
	}

	return products, nil
}
