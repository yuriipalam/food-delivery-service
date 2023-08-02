package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/service"
	"food_delivery/utils"
	"net/http"
)

type ProductHandler struct {
	repo repository.ProductRepositoryI
	service *service.ProductService
}

func NewProductHandler(repo repository.ProductRepositoryI, service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		repo: repo,
		service: service,
	}
}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	productResp, err := ph.service.GetProductResponseByID(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	} else if productResp == nil {
		response.SendNotFoundError(w, fmt.Errorf("product not found"))
		return
	}

	response.SendOK(w, productResp)
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	productResponses, err := ph.service.GetAllProductResponses()
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(productResponses) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	response.SendOK(w, productResponses)
}

func (ph *ProductHandler) GetAllProductsBySupplierID(w http.ResponseWriter, r *http.Request) {
	supplierID, err := utils.GetIntValueByKeyFromMuxVars("supplier_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	productResponses, err := ph.service.GetProductResponsesBySupplierID(supplierID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(productResponses) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	response.SendOK(w, productResponses)
}

func (ph *ProductHandler) GetAllProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetIntValueByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	productResponses, err := ph.service.GetProductResponsesByCategoryID(categoryID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(productResponses) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	response.SendOK(w, productResponses)
}

func (ph *ProductHandler) GetAllProductsBySupplierIDAndCategoryID(w http.ResponseWriter, r *http.Request) {
	supplierID, err := utils.GetIntValueByKeyFromMuxVars("supplier_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	categoryID, err := utils.GetIntValueByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	productResponses, err := ph.service.GetProductResponsesBySupplierIDAndCategoryID(supplierID, categoryID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if len(productResponses) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	response.SendOK(w, productResponses)
}
