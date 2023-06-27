package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
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
