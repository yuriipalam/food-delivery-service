package handler

import (
	"fmt"
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
		response.SendBadRequestError(w, err)
		return
	}

	category, err := ch.repo.GetCategoryByID(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	} else if category == nil {
		response.SendNotFoundError(w, fmt.Errorf("category with id %d not found", id))
		return
	}

	response.SendOK(w, category)
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

	response.SendOK(w, categories)
}
