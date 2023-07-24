package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/utils"
	"net/http"
)

type CategoryHandler struct {
	repo repository.CategoryRepositoryI
}

func NewCategoryHandler(repo repository.CategoryRepositoryI) *CategoryHandler {
	return &CategoryHandler{
		repo: repo,
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

	res, err := ch.getCategoryResponseFromModel(category)
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

	res, err := ch.getCategoryResponsesFromModels(categories)
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

	res, err := ch.getCategoryResponsesFromModels(categories)
	if err != nil {
		response.SendInternalServerError(w, err)
	}

	response.SendOK(w, res)
}

func (ch *CategoryHandler) getCategoryResponseFromModel(category *model.Category) (*response.CategoryResponse, error) {
	var categoryRes response.CategoryResponse

	categoryMarshaled, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(categoryMarshaled, &categoryRes); err != nil {
		return nil, fmt.Errorf("cannot unmarshal category from db into response")
	}

	categoryRes.URL = fmt.Sprintf("/categories/%d", category.ID)
	categoryRes.ImageURL = fmt.Sprintf("%s/images/categories/%d/%s", config.Root, category.ID, category.Image)

	return &categoryRes, nil
}

func (ch *CategoryHandler) getCategoryResponsesFromModels(categories []model.Category) ([]response.CategoryResponse, error) {
	var categoriesRes []response.CategoryResponse

	for _, categoryRes := range categories {
		categoryRes, err := ch.getCategoryResponseFromModel(&categoryRes)
		if err != nil {
			return nil, err
		}

		categoriesRes = append(categoriesRes, *categoryRes)
	}

	return categoriesRes, nil
}
