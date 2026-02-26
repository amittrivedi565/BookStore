package repository

import (
	"bookstore/models"
	"database/sql"
	"fmt"
)

type DB struct {
	DB *sql.DB
}

func InjectDB(db *sql.DB) *DB {
	return &DB{
		DB: db,
	}
}

func (db *DB) CreateBook(book *models.Book) (int64, error) {

	query := `
	INSERT INTO books 
	(title, description, author_id, category_id, price, stock, isbn)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := db.DB.Exec(
		query,
		book.Title,
		book.Description,
		book.AuthorId,
		book.CategoryId,
		book.Price,
		book.Stock,
		book.Isbn,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (db *DB) GetBooks() ([]models.Book, error) {

	query := `
	SELECT id, title, description, author_id, category_id, price, stock, isbn
	FROM books
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.AuthorId,
			&book.CategoryId,
			&book.Price,
			&book.Stock,
			&book.Isbn,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (db *DB) GetBookById(bookId int) (*models.BookResponse, error) {

	query := `
	SELECT 
		b.id,
		b.title,
		b.description,
		b.price,
		b.stock,
		b.isbn,

		a.id,
		a.name,
		a.description,

		c.id,
		c.name,
		c.description

	FROM books b
	JOIN authors a ON b.author_id = a.id
	JOIN categories c ON b.category_id = c.id
	WHERE b.id = ?
	`

	var book models.BookResponse

	err := db.DB.QueryRow(query, bookId).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.Price,
		&book.Stock,
		&book.ISBN,

		&book.Author.ID,
		&book.Author.Name,
		&book.Author.Description,

		&book.Category.ID,
		&book.Category.Name,
		&book.Category.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	return &book, nil
}

func (db *DB) UpdateBook(bookId int, book *models.Book) error {

	query := `
	UPDATE books
	SET title = ?, 
	    description = ?, 
	    author_id = ?, 
	    category_id = ?, 
	    price = ?, 
	    stock = ?, 
	    isbn = ?
	WHERE id = ?
	`

	result, err := db.DB.Exec(
		query,
		book.Title,
		book.Description,
		book.AuthorId,
		book.CategoryId,
		book.Price,
		book.Stock,
		book.Isbn,
		bookId,
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

func (db *DB) DeleteBook(bookId int) error {

	query := `DELETE FROM books WHERE id = ?`

	result, err := db.DB.Exec(query, bookId)
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

func (db *DB) GetBooksByAuthorName(authorName string) ([]models.Book, error) {

	query := `
	SELECT 
		b.id, 
		b.title, 
		b.description,
		b.author_id,
		b.category_id,
		b.price,
		b.stock,
		b.isbn
	FROM books b
	JOIN authors a ON b.author_id = a.id
	WHERE a.name = ?
	`

	rows, err := db.DB.Query(query, authorName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.AuthorId,
			&book.CategoryId,
			&book.Price,
			&book.Stock,
			&book.Isbn,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (db *DB) GetBooksByCategory(categoryName string) ([]models.BookResponse, error) {

	query := `
	SELECT 
		b.id,
		b.title,
		b.description,
		a.name,
		b.price,
		b.stock,
		b.isbn
	FROM books b
	JOIN authors a ON b.author_id = a.id
	JOIN categories c ON b.category_id = c.id
	WHERE c.name = ?
	`

	rows, err := db.DB.Query(query, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookResponse

	for rows.Next() {
		var book models.BookResponse

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Author,
			&book.Price,
			&book.Stock,
			&book.ISBN,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, rows.Err()
}