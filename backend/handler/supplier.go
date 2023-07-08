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
		response.SendBadRequestError(w, err)
		return
	}

	supplier, err := sh.repo.GetSupplierByID(id)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if supplier == nil {
		response.SendNotFoundError(w, fmt.Errorf("cannot find supplier with id %d", id))
		return
	}

	supplierRes, err := sh.GetSupplierResponseFromModel(supplier)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) GetSuppliersByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetIntValueByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	suppliers, err := sh.repo.GetSuppliersByCategoryID(categoryID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(suppliers) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no suppliers found with category_id %d", categoryID))
		return
	}

	supplierRes, err := sh.GetSuppliersResponseFromModel(suppliers)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := sh.repo.GetAllSuppliers()
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(suppliers) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no suppliers found"))
		return
	}

	supplierRes, err := sh.GetSuppliersResponseFromModel(suppliers)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) GetSupplierResponseFromModel(supplier *model.Supplier) (*response.SupplierResponse, error) {
	var supplierRes response.SupplierResponse

	supplierMarshaled, err := json.Marshal(supplier)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal supplier from db")
	}

	err = json.Unmarshal(supplierMarshaled, &supplierRes)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal supplier from db into response")
	}

	supplierRes.CategoryName, err = sh.repo.GetCategoryNameByID(supplier.CategoryID)
	if err != nil {
		return nil, err
	}

	return &supplierRes, nil
}

func (sh *SupplierHandler) GetSuppliersResponseFromModel(suppliers []model.Supplier) ([]response.SupplierResponse, error) {
	var suppliersRes []response.SupplierResponse

	for _, supplier := range suppliers {
		supplierRes, err := sh.GetSupplierResponseFromModel(&supplier)
		if err != nil {
			return nil, err
		}

		suppliersRes = append(suppliersRes, *supplierRes)
	}

	return suppliersRes, nil
}
