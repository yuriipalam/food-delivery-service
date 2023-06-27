package handler

import (
	"encoding/json"
	"food_delivery/repository"
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

func (sh *SupplierHandler) GetListOfSuppliers(w http.ResponseWriter, r *http.Request) {
	// call db and get list

	suppliers, err := sh.repo.GetAllSuppliers()
	if err != nil {
		// prompt error
	}

	data, err := json.Marshal(suppliers)
	if err != nil {
		// prompt error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
