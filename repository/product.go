package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"food_delivery/model"
	"strings"
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
		var driverValue driver.Value

		err := rows.Scan(
			&product.ID,
			&product.SupplierID,
			&product.Name,
			&product.Image,
			&product.Description,
			&driverValue,
			&product.Price,
		)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("cannot scan product")
		}

		strSlice, err := convertDriverValueToStrSlice(driverValue)
		if err != nil {
			return nil, err
		}
		product.Ingredients = strSlice

		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("type mismatch on scanning rows")
	}

	return products, nil
}

func convertStrSliceToSqlArr(strSlice []string) string {
	for i, str := range strSlice {
		strSlice[i] = "'" + str + "'"
	}

	return "{" + strings.Join(strSlice, ",") + "}"
}

func convertDriverValueToStrSlice(value driver.Value) ([]string, error) {
	bytesValue, ok := value.([]byte)
	if !ok {
		return nil, fmt.Errorf("error converting driver.Value to []bytes")
	}

	return strings.Split(strings.Trim(string(bytesValue), "{}"), ","), nil
}
