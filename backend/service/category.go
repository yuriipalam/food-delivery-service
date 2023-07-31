package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
	"food_delivery/response"
)

type CategoryService struct {
	cfg *config.Config
	db  *sql.DB
}

func NewCategoryService(cfg *config.Config, db *sql.DB) *CategoryService {
	return &CategoryService{
		cfg: cfg,
		db:  db,
	}
}

func (cs *CategoryService) GetCategoryResponseFromModel(category *model.Category) (*response.CategoryResponse, error) {
	var categoryRes response.CategoryResponse

	categoryMarshaled, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(categoryMarshaled, &categoryRes); err != nil {
		return nil, fmt.Errorf("cannot unmarshal category from db into response")
	}

	categoryRes.URL = fmt.Sprintf("/categories/%d", category.ID)
	// prod
	categoryRes.ImageURL = fmt.Sprintf("/images/categories/%d/%s", category.ID, category.Image)

	// dev
	// categoryRes.ImageURL = fmt.Sprintf("%s/images/categories/%d/%s", config.Root, category.ID, category.Image)

	return &categoryRes, nil
}

func (cs *CategoryService) GetCategoryResponsesFromModels(categories []model.Category) ([]response.CategoryResponse, error) {
	var categoryResponses []response.CategoryResponse

	for _, categoryRes := range categories {
		categoryRes, err := cs.GetCategoryResponseFromModel(&categoryRes)
		if err != nil {
			return nil, err
		}

		categoryResponses = append(categoryResponses, *categoryRes)
	}

	return categoryResponses, nil
}
