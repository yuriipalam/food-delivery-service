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

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("id must be integer"))
		return
	}

	product, err := ph.repo.GetProductByID(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	if product == nil {
		response.SendNotFoundError(w, fmt.Errorf("no products found with given id %d", id))
		return
	}

	response.SendOK(w, product)
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
		response.SendBadRequestError(w, fmt.Errorf("supplier_id must be integer"))
		return
	}

	products, err := ph.repo.GetAllProductsBySupplierID(supplierID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products with supplier_id %d found", supplierID))
		return
	}

	response.SendOK(w, products)
}

func (ph *ProductHandler) GetAllProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("category_id must be integer"))
		return
	}

	products, err := ph.repo.GetAllProductsByCategoryID(categoryID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products with category_id %d found", categoryID))
		return
	}

	response.SendOK(w, products)
}

func (ph *ProductHandler) GetAllProductsBySupplierIDAndCategoryID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplierID, err := strconv.Atoi(vars["supplier_id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("supplier_id must be integer"))
		return
	}
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("category_id must be integer"))
		return
	}

	products, err := ph.repo.GetAllProductsBySupplierIDAndCategoryID(supplierID, categoryID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products with supplier_id %d and category_id %d found", supplierID, categoryID))
		return
	}

	response.SendOK(w, products)
}
