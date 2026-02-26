package repository

import (
	"bookstore/models"
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func InjectDB(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateBook(book models.Book) error {
	query := `INSERT INTO books (title, author, price, stock isbn) VALUES (?,?,?,?,?)`

	_, err := r.DB.Exec(query,
		book.Title,
		book.Author,
		book.Price,
		book.Stock,
		book.Isbn,
	)

	return err
}

func (r *Repository) GetBooks() ([]models.Book, error) {
	query := `SELECT id, title, author, price, stock, isbn FROM books`

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Price,
			&book.Stock,
			&book.Isbn,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	defer rows.Close()

	return books, nil
}

func (r *Repository) GetBookById(bookId int) (*models.Book, error) {
	query := `SELECT id, title, author, price, stock, isbn 
	FROM books 
	WHERE id = ? `

	var book models.Book

	err := r.DB.QueryRow(query, bookId).Scan(
		&book.Title,
		&book.Author,
		&book.Price,
		&book.Stock,
		&book.Isbn,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}
	return &book, err
}

func (r *Repository) UpdateBookById(id int, book *models.Book) error {
	query := `UPDATE books SET title = ?, author = ?, price = ?, stock = ?, isbn = ? WHERE id = ?`

	result, err := r.DB.Exec(query,
		book.Title,
		book.Author,
		book.Price,
		book.Stock,
		book.Isbn,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}

func (r *Repository) DeleteBookByID(id int) error {

	query := `DELETE FROM books WHERE id = ?`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}
