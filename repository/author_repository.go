package repository

import (
	"bookstore/models"
	"database/sql"
	"fmt"
)


func (db *DB) GetAuthors() ([]models.Author, error) {

	query := `SELECT id, name, description FROM authors`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author

	for rows.Next() {
		var author models.Author

		if err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Description,
		); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, rows.Err()
}

func (db *DB) GetAuthorByID(id int) (*models.Author, error) {

	query := `SELECT id, name, description FROM authors WHERE id = ?`

	var author models.Author

	err := db.DB.QueryRow(query, id).Scan(
		&author.ID,
		&author.Name,
		&author.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("author not found")
		}
		return nil, err
	}

	return &author, nil
}

func (db *DB) CreateAuthor(author *models.Author) (int64, error) {

	query := `INSERT INTO authors (name, description) VALUES (?, ?)`

	res, err := db.DB.Exec(
		query,
		author.Name,
		author.Description,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (db *DB) UpdateAuthor(id int, author *models.Author) error {

	query := `
	UPDATE authors
	SET name = ?, description = ?
	WHERE id = ?
	`

	res, err := db.DB.Exec(
		query,
		author.Name,
		author.Description,
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
		return fmt.Errorf("author not found")
	}

	return nil
}

func (db *DB) DeleteAuthor(id int) error {

	query := `DELETE FROM authors WHERE id = ?`

	res, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("author not found")
	}

	return nil
}