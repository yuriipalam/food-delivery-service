package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/service"
	"food_delivery/utils"
	"net/http"
)

type SupplierHandler struct {
	repo repository.SupplierRepositoryI
	service *service.SupplierService
}

func NewSupplierHandler(repo repository.SupplierRepositoryI, service *service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		repo: repo,
		service: service,
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
		response.SendInternalServerError(w, err)
		return
	} else if supplier == nil {
		response.SendNotFoundError(w, fmt.Errorf("supplier not found"))
		return
	}

	supplierRes, err := sh.service.GetSupplierResponseBySupplierID(supplier.ID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, supplierRes)
}

func (sh *SupplierHandler) GetSuppliersByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryIDs, err := utils.GetIntSliceByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integers"
		return
	}

	suppliers, err := sh.repo.GetSuppliersByCategoryIDs(categoryIDs)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(suppliers) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no suppliers found"))
		return
	}

	supplierRes, err := sh.service.GetSupplierResponses(suppliers)
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

	supplierRes, err := sh.service.GetSupplierResponses(suppliers)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, supplierRes)
}
