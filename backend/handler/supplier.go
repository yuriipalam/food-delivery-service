package handler

import (
	"fmt"
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

	response.SendOK(w, supplier)
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

	response.SendOK(w, suppliers)
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

	response.SendOK(w, suppliers)
}
