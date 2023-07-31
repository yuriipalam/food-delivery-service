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

func (ss *SupplierService) GetSupplierResponseFromModel(supplier *model.Supplier) (*response.SupplierResponse, error) {
	query := `SELECT supplier_id, category_id, c.name
    		  FROM supplier_category ss
    		  JOIN category c ON ss.category_id = c.id
    		  WHERE supplier_id = $1`

	stmt, err := ss.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query(supplier.ID)
	if err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	supplierCategoriesDict, err := ss.getSupplierCategoriesDictFromRows(rows)
	if err != nil {
		return nil, err
	}

	supplierResponse := response.SupplierResponse{
		ID:          supplier.ID,
		Categories:  supplierCategoriesDict[supplier.ID],
		Name:        supplier.Name,
		Description: supplier.Description,
		TimeOpening: supplier.TimeOpening,
		TimeClosing: supplier.TimeClosing,
	}

	supplierResponse.URL = fmt.Sprintf("/suppliers/%d", supplier.ID)

	// production
	supplierResponse.ImageURL = fmt.Sprintf("/images/suppliers/%d/%s", supplier.ID, supplier.Image)

	// development
	// supplierResponse.ImageURL = fmt.Sprintf("%s/images/suppliers/%d/%s", config.Root, supplier.ID, supplier.Image)

	return &supplierResponse, nil
}

func (ss *SupplierService) GetSupplierResponses(suppliers []model.Supplier) ([]response.SupplierResponse, error) {
	query := "SELECT supplier_id, category_id, c.name FROM supplier_category ss JOIN category c ON ss.category_id = c.id"

	stmt, err := ss.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	supplierCategoriesDict, err := ss.getSupplierCategoriesDictFromRows(rows)
	if err != nil {
		return nil, err
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
		// prod
		supplierResponse.ImageURL = fmt.Sprintf("/images/suppliers/%d/%s", supplier.ID, supplier.Image)

		// dev
		// supplierResponse.ImageURL = fmt.Sprintf("%s/images/suppliers/%d/%s", config.Root, supplier.ID, supplier.Image)

		supplierResponses = append(supplierResponses, supplierResponse)
	}

	return supplierResponses, nil
}

func (ss *SupplierService) getSupplierCategoriesDictFromRows(rows *sql.Rows) (map[int][]response.SupplierCategoryResponse, error) {
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

	return supplierCategoriesDict, nil
}