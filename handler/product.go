package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"net/http"
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
