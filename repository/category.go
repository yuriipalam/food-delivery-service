package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type CategoryRepositoryI interface {
	GetAllCategories() ([]model.Category, error)
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM category")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("cannot execute query")
	}

	var categories []model.Category

	for rows.Next() {
		var category model.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Image,
			&category.Description,
			)
		if err != nil {
			return nil, fmt.Errorf("cannot scan category")
		}

		categories = append(categories, category)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("type mismatch on scanning rows")
	}

	return categories, nil
}
