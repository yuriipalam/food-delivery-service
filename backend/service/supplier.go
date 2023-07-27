package service

import (
	"database/sql"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
	"food_delivery/response"
)

type SupplierService struct {
	cfg *config.Config
	db  *sql.DB
}

func NewSupplierService(cfg *config.Config, db *sql.DB) *SupplierService {
	return &SupplierService{
		cfg: cfg,
		db:  db,
	}
}

func (sc *SupplierService) GetSupplierResponseBySupplierID(id int) (*response.SupplierResponse, error) {
	query := `SELECT s.id, s.name, s.image, s.description, s.time_opening, s.time_closing, category_id, c.name
			  FROM supplier s
			  JOIN supplier_category sc ON s.id = sc.supplier_id
			  JOIN category c ON sc.category_id = c.id
			  WHERE s.id = $1`

	stmt, err := sc.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for supplier_id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("cannot run query for supplier_id %d", id)
	}

	var supplierResponse response.SupplierResponse
	var image string
	var categories []response.SupplierCategoryResponse

	for rows.Next() {
		var category response.SupplierCategoryResponse

		err := rows.Scan(
			&supplierResponse.ID,
			&supplierResponse.Name,
			&image,
			&supplierResponse.Description,
			&supplierResponse.TimeOpening,
			&supplierResponse.TimeClosing,
			&category.CategoryID,
			&category.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan supplier %d", id)
		}

		categories = append(categories, category)
	}

	supplierResponse.URL = fmt.Sprintf("/suppliers/%d", id)
	supplierResponse.ImageURL = fmt.Sprintf("%s/images/suppliers/%d/%s", config.Root, id, image)

	return &supplierResponse, nil
}

func (sc *SupplierService) GetSupplierResponses(suppliers []model.Supplier) ([]response.SupplierResponse, error) {
	query := "SELECT supplier_id, category_id, c.name FROM supplier_category sc JOIN category c ON sc.category_id = c.id"

	stmt, err := sc.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	supplierCategoriesDict := make(map[int][]response.SupplierCategoryResponse)

	for rows.Next() {
		var supID int
		var catID int
		var catName string

		err := rows.Scan(
			&supID,
			&catID,
			&catName,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan")
		}

		supCatResponse := response.SupplierCategoryResponse{
			CategoryID:   catID,
			CategoryName: catName,
		}

		if supplierCategoriesDict[supID] == nil {
			supplierCategoriesDict[supID] = []response.SupplierCategoryResponse{supCatResponse}
		} else {
			supplierCategoriesDict[supID] = append(supplierCategoriesDict[supID], supCatResponse)
		}
	}

	var supplierResponses []response.SupplierResponse

	for _, supplier := range suppliers {
		supplierResponse := response.SupplierResponse{
			ID:          supplier.ID,
			Categories:  supplierCategoriesDict[supplier.ID],
			Name:        supplier.Name,
			Description: supplier.Description,
			TimeOpening: supplier.TimeOpening,
			TimeClosing: supplier.TimeClosing,
		}

		supplierResponse.URL = fmt.Sprintf("/suppliers/%d", supplier.ID)
		supplierResponse.ImageURL = fmt.Sprintf("%s/images/suppliers/%d/%s", config.Root, supplier.ID, supplier.Image)

		supplierResponses = append(supplierResponses, supplierResponse)
	}

	return supplierResponses, nil
}
