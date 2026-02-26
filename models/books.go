package models

type Book struct {
	ID          int     `json:"id"`
	AuthorId    string  `json:"author_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CategoryId    int     `json:"category_id"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Isbn        string  `json:"isbn"`
}

type BookResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	ISBN        string   `json:"isbn"`
	Author      AuthorResponse   `json:"author"`
	Category    CategoryResponse `json:"category"`
}

type AuthorResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}