package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("id must be integer"))
		return
	}

	supplier, err := sh.repo.GetSupplierByID(id)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if supplier == nil {
		response.SendNotFoundError(w, fmt.Errorf("cannot find supplier with id %d", id))
		return
	}

	response.SendOK(w, supplier)
}

func (sh *SupplierHandler) GetSupplierByCategoryID(w http.ResponseWriter, r *http.Request) {

}

func (sh *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := sh.repo.GetAllSuppliers()
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(suppliers) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no suppliers found"))
		return
	}

	response.SendOK(w, suppliers)
}
