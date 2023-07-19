package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/utils"
	"net/http"
)

type SupplierHandler struct {
	repo repository.SupplierRepositoryI
}

func NewSupplierHandler(repo repository.SupplierRepositoryI) *SupplierHandler {
	return &SupplierHandler{
		repo: repo,
	}
}

func (sh *SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	supplier, err := sh.repo.GetSupplierByID(id)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch supplier"))
		return
	} else if supplier == nil {
		response.SendNotFoundError(w, fmt.Errorf("supplier not found"))
		return
	}

	supplierRes, err := sh.getSupplierResponseFromModel(supplier)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create response"))
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) GetSuppliersByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryIDs, err := utils.GetIntSliceByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "ids must be integers"
		return
	}

	suppliers, err := sh.repo.GetSuppliersByCategoryIDs(categoryIDs)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch suppliers"))
		return
	} else if len(suppliers) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no suppliers found"))
		return
	}

	supplierRes, err := sh.getSupplierResponsesFromModel(suppliers)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create responses"))
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := sh.repo.GetAllSuppliers()
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch suppliers"))
		return
	} else if len(suppliers) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no suppliers found"))
		return
	}

	supplierRes, err := sh.getSupplierResponsesFromModel(suppliers)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create responses"))
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) getSupplierResponseFromModel(supplier *model.Supplier) (*response.SupplierResponse, error) {
	var supplierRes response.SupplierResponse

	supplierMarshaled, err := json.Marshal(supplier)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal supplier from db")
	}

	err = json.Unmarshal(supplierMarshaled, &supplierRes)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal supplier from db into response")
	}

	supplierRes.URL = fmt.Sprintf("/suppliers/%d", supplier.ID)
	supplierRes.ImageURL = fmt.Sprintf("http://localhost:8080/images/suppliers/%s", supplier.Image)

	supplierRes.Categories, err = sh.repo.GetCategoryResponsesBySupplierID(supplier.ID)
	if err != nil {
		return nil, err
	}

	return &supplierRes, nil
}

func (sh *SupplierHandler) getSupplierResponsesFromModel(suppliers []model.Supplier) ([]response.SupplierResponse, error) {
	var suppliersRes []response.SupplierResponse

	for _, supplier := range suppliers {
		supplierRes, err := sh.getSupplierResponseFromModel(&supplier)
		if err != nil {
			return nil, err
		}

		suppliersRes = append(suppliersRes, *supplierRes)
	}

	return suppliersRes, nil
}
