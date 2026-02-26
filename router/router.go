package routes

import (
	"bookstore/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *controller.Handler) {

	// =========================
	// BOOK ROUTES
	// =========================

	books := r.Group("/books")
	{
		books.POST("", h.CreateBook)
		books.GET("", h.GetBooks)
		books.GET("/:id", h.GetBookByID)
		books.PUT("/:id", h.UpdateBook)
		books.DELETE("/:id", h.DeleteBook)

		books.GET("/author", h.GetBooksByAuthorName)
		books.GET("/category", h.GetBooksByCategory)
	}

	// =========================
	// AUTHOR ROUTES
	// =========================

	authors := r.Group("/authors")
	{
		authors.GET("", h.GetAuthors)
		authors.GET("/:id", h.GetAuthorByID)
		authors.POST("", h.CreateAuthor)
		authors.PUT("/:id", h.UpdateAuthor)
		authors.DELETE("/:id", h.DeleteAuthor)
	}

	// =========================
	// CATEGORY ROUTES
	// =========================

	categories := r.Group("/categories")
	{
		categories.GET("", h.GetCategories)
		categories.GET("/:id", h.GetCategoryByID)
		categories.POST("", h.CreateCategory)
		categories.PUT("/:id", h.UpdateCategory)
		categories.DELETE("/:id", h.DeleteCategory)
	}
}