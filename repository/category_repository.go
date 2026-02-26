package repository

import (
	"bookstore/models"
	"database/sql"
	"fmt"
)

func (db *DB) GetCategories() ([]models.Category, error) {

	query := `SELECT id, name, description FROM categories`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category

		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, rows.Err()
}

func (db *DB) GetCategoryByID(id int) (*models.Category, error) {

	query := `SELECT id, name, description FROM categories WHERE id = ?`

	var category models.Category

	err := db.DB.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, err
	}

	return &category, nil
}

func (db *DB) CreateCategory(category *models.Category) (int64, error) {

	query := `INSERT INTO categories (name, description) VALUES (?, ?)`

	res, err := db.DB.Exec(
		query,
		category.Name,
		category.Description,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (db *DB) UpdateCategory(id int, category *models.Category) error {

	query := `
	UPDATE categories
	SET name = ?, description = ?
	WHERE id = ?
	`

	res, err := db.DB.Exec(
		query,
		category.Name,
		category.Description,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

func (db *DB) DeleteCategory(id int) error {

	query := `DELETE FROM categories WHERE id = ?`

	res, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}
