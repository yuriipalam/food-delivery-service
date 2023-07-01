package handler

import (
	"fmt"
	"food_delivery/repository"
	"food_delivery/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("id must be integer"))
		return
	}

	category, err := ch.repo.GetCategoryByID(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	if category == nil {
		response.SendNotFoundError(w, fmt.Errorf("no category found with id %d", id))
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
