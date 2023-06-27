package handler

import (
	"food_delivery/repository"
	"food_delivery/response"
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

func (sh *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := sh.repo.GetAllSuppliers()
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, suppliers)
}
