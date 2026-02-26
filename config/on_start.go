package config

import "database/sql"

/* On start of application create these tables*/
type db struct {
	DB *sql.DB
}

func inject_db_instance(instance *sql.DB) *db {
	return &db{
		DB: instance,
	}
}

func (s *db) on_start() {

	create_book_table_query := `
	CREATE TABLE IF NOT EXISTS books (
	id INT AUTO_INCREMENT PRIMARY KEY,
	author_id INT,
	category_id INT,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	price DECIMAL(10,2),
	stock INT DEFAULT 0,
	isbn VARCHAR(20) UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (author_id) REFERENCES authors(id),
	FOREIGN KEY (category_id) REFERENCES categories(id)
	);`

	s.DB.Exec(create_book_table_query)

	create_author_table_query := `
	CREATE TABLE IF NOT EXISTS authors (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(150) NOT NULL,
	about TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	s.DB.Exec(create_author_table_query)
}
