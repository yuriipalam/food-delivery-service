package repository

import (
	"database/sql"
	"fmt"
	"food_delivery/model"
)

type CategoryRepositoryI interface {
	GetCategoryByID(int) (*model.Category, error)
	GetCategoriesBySupplierID(int) ([]model.Category, error)
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

func (cr *CategoryRepository) GetCategoryByID(id int) (*model.Category, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM category WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for id %d", id)
	}

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, fmt.Errorf("cannot run query for category id %d", id)
	}

	var category model.Category

	err = row.Scan(
		&category.ID,
		&category.Name,
		&category.Image,
		&category.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot scan category with id %d", id)
	}

	return &category, nil
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

	return categories, nil
}

func (cr *CategoryRepository) GetCategoriesBySupplierID(id int) ([]model.Category, error) {
	stmt, err := cr.db.Prepare("SELECT * FROM category WHERE id IN (SELECT category_id FROM supplier WHERE id = $1)")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement for supplier id %d", id)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("cannot run query for supplier id %d", id)
	}

	var categories []model.Category

	for rows.Next() {
		var category model.Category

		err = rows.Scan(
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

	return categories, nil
}
