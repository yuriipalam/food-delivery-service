package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	repo repository.ProductRepositoryI
}

func NewProductHandler(repo repository.ProductRepositoryI) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.repo.GetAllProducts()
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	response.SendOK(w, products)
}

func (ph *ProductHandler) GetAllProductsBySupplierID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplierID, err := strconv.Atoi(vars["supplier_id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("supplied_id must be integer"))
	}

	products, err := ph.repo.GetAllProductsBySupplierID(supplierID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products with supplied_id %d found", supplierID))
		return
	}

	response.SendOK(w, products)
}
