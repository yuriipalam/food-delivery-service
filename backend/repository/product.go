package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"food_delivery/model"
	"strings"
)

type ProductRepositoryI interface {
	GetProductByID(int) (*model.Product, error)
	GetAllProducts() ([]model.Product, error)
	GetAllProductsBySupplierID(int) ([]model.Product, error)
	GetAllProductsByCategoryID(int) ([]model.Product, error)
	GetAllProductsBySupplierIDAndCategoryID(int, int) ([]model.Product, error)
	GetSupplierNameByID(int) (string, error)
	GetCategoryNameByID(int) (string, error)
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
	var driverValue driver.Value

	err = row.Scan(
		&product.ID,
		&product.SupplierID,
		&product.CategoryID,
		&product.Name,
		&product.Image,
		&product.Description,
		&driverValue,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan product")
	}

	strSlice, err := convertDriverValueToStrSlice(driverValue)
	if err != nil {
		return nil, err
	}

	product.Ingredients = strSlice

	return &product, nil
}

func (pr *ProductRepository) GetAllProducts() ([]model.Product, error) {
	return pr.selectProductsQuery("SELECT * FROM product")
}

func (pr *ProductRepository) GetAllProductsBySupplierID(id int) ([]model.Product, error) {
	return pr.selectProductsQuery("SELECT * FROM product WHERE supplier_id = $1", id)
}

func (pr *ProductRepository) GetAllProductsByCategoryID(id int) ([]model.Product, error) {
	return pr.selectProductsQuery("SELECT * FROM product WHERE category_id = $1", id)
}

func (pr *ProductRepository) GetAllProductsBySupplierIDAndCategoryID(sID int, cID int) ([]model.Product, error) {
	return pr.selectProductsQuery("SELECT * FROM product WHERE supplier_id = $1 AND category_id = $2", sID, cID)
}

func (pr *ProductRepository) GetSupplierNameByID(id int) (string, error) {
	stmt, err := pr.db.Prepare("SELECT name FROM supplier WHERE id = $1")
	if err != nil {
		return "", fmt.Errorf("cannot prepare statement for supplier id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return "", fmt.Errorf("cannot run query for supplier id %d", id)
	}

	var name string

	if err := row.Scan(&name); err != nil {
		return "", fmt.Errorf("cannot scan supplier name")
	}

	return name, nil
}

func (pr *ProductRepository) GetCategoryNameByID(id int) (string, error) {
	stmt, err := pr.db.Prepare("SELECT name FROM category WHERE id = $1")
	if err != nil {
		return "", fmt.Errorf("cannot prepare statement for category id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return "", fmt.Errorf("cannot run query for category id %d", id)
	}

	var name string

	if err := row.Scan(&name); err != nil {
		return "", fmt.Errorf("cannot scan category name")
	}

	return name, nil
}

func (pr *ProductRepository) selectProductsQuery(query string, data ...any) ([]model.Product, error) {
	stmt, err := pr.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query(data...)
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
			&product.CategoryID,
			&product.Name,
			&product.Image,
			&product.Description,
			&driverValue,
			&product.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan product")
		}

		strSlice, err := convertDriverValueToStrSlice(driverValue)
		if err != nil {
			return nil, err
		}
		product.Ingredients = strSlice

		products = append(products, product)
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

	return strings.Split(strings.Trim(strings.ReplaceAll(string(bytesValue), "\"", ""), "{}"), ","), nil
}
