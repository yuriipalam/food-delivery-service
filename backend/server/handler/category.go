package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/service"
	"food_delivery/utils"
	"net/http"
)

type CategoryHandler struct {
	repo repository.CategoryRepositoryI
	service *service.CategoryService
}

func NewCategoryHandler(repo repository.CategoryRepositoryI, service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		repo: repo,
		service: service,
	}
}

func (ch *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	category, err := ch.repo.GetCategoryByID(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	} else if category == nil {
		response.SendNotFoundError(w, fmt.Errorf("category not found"))
		return
	}

	res, err := ch.service.GetCategoryResponseFromModel(category)
	if err != nil {
		response.SendInternalServerError(w, err)
	}

	response.SendOK(w, res)
}

func (ch *CategoryHandler) GetCategoriesBySupplierID(w http.ResponseWriter, r *http.Request) {
	supplierID, err := utils.GetIntValueByKeyFromMuxVars("supplier_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be int"
		return
	}

	categories, err := ch.repo.GetCategoriesBySupplierID(supplierID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(categories) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("categories not found"))
		return
	}

	res, err := ch.service.GetCategoryResponsesFromModels(categories)
	if err != nil {
		response.SendInternalServerError(w, err)
	}

	response.SendOK(w, res)
}

func (ch *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := ch.repo.GetAllCategories()
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(categories) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no categories found"))
		return
	}

	res, err := ch.service.GetCategoryResponsesFromModels(categories)
	if err != nil {
		response.SendInternalServerError(w, err)
	}

	response.SendOK(w, res)
}
