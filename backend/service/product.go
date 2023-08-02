package service

import (
	"database/sql"
	"fmt"
	"food_delivery/config"
	"food_delivery/response"
)

type ProductService struct {
	cfg *config.Config
	db  *sql.DB
}

func NewProductService(cfg *config.Config, db *sql.DB) *ProductService {
	return &ProductService{
		cfg: cfg,
		db:  db,
	}
}

func (ps *ProductService) GetProductResponseByID(id int) (*response.ProductResponse, error) {
	query := `SELECT id, supplier_id, supplier_name, category_id, category_name, product.name, image, description, price
			  FROM product,
				   (SELECT s.name FROM supplier s, product p WHERE s.id = p.supplier_id) AS supplier_name,
				   (SELECT c.name FROM category c, product p WHERE c.id = p.category_id) AS category_name
			  WHERE product.id = $1`

	stmt, err := ps.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	row := stmt.QueryRow(id)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	var productResp response.ProductResponse
	var imageName string

	err = row.Scan(
		&productResp.ID,
		&productResp.SupplierID,
		&productResp.SupplierName,
		&productResp.CategoryID,
		&productResp.CategoryName,
		&productResp.Name,
		&imageName,
		&productResp.Description,
		&productResp.Price,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot scan row")
	}

	productResp.URL = fmt.Sprintf("/products/%d", id)
	productResp.ImageURL = fmt.Sprintf("/images/products/%d/%s", productResp.SupplierID, imageName)

	return &productResp, nil
}

func (ps *ProductService) GetAllProductResponses() ([]response.ProductResponse, error) {
	return ps.getProductResponsesByCondition("")
}

func (ps *ProductService) GetProductResponsesBySupplierID(id int) ([]response.ProductResponse, error) {
	condition := "WHERE supplier_id = $1"
	return ps.getProductResponsesByCondition(condition, id)
}

func (ps *ProductService) GetProductResponsesByCategoryID(id int) ([]response.ProductResponse, error) {
	condition := "WHERE category_id = $1"
	return ps.getProductResponsesByCondition(condition, id)
}

func (ps *ProductService) GetProductResponsesBySupplierIDAndCategoryID(sID int, cID int) ([]response.ProductResponse, error) {
	condition := "WHERE supplier_id = $1 AND category_id = $2"
	return ps.getProductResponsesByCondition(condition, sID, cID)
}

func (ps *ProductService) getProductResponsesByCondition(condition string, ids ...any) ([]response.ProductResponse, error) {
	query := fmt.Sprintf(
		`SELECT p.id, supplier_id, s.name, category_id, s.name, p.name, p.image, p.description, price
				FROM product p
				JOIN supplier s ON p.supplier_id = s.id
				JOIN category c ON p.category_id = c.id
				%s`,
		condition,
	)

	stmt, err := ps.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query(ids...)
	if err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	var productResponses []response.ProductResponse

	for rows.Next() {
		var productResp response.ProductResponse
		var imageName string

		err = rows.Scan(
			&productResp.ID,
			&productResp.SupplierID,
			&productResp.SupplierName,
			&productResp.CategoryID,
			&productResp.CategoryName,
			&productResp.Name,
			&imageName,
			&productResp.Description,
			&productResp.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan row")
		}

		productResp.URL = fmt.Sprintf("/products/%d", productResp.ID)
		productResp.ImageURL = fmt.Sprintf("/images/products/%d/%s", productResp.SupplierID, imageName)

		productResponses = append(productResponses, productResp)
	}

	return productResponses, nil
}
