package controller

import (
	"bookstore/models"
	"bookstore/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *repository.DB
}

func NewHandler(repo *repository.DB) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateBook(c *gin.Context) {

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "book created successfully",
		"id":      id,
	})
}

func (h *Handler) GetBooks(c *gin.Context) {

	books, err := h.repo.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBookByID(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	book, err := h.repo.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) UpdateBook(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateBook(id, &book); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated successfully",
	})
}

func (h *Handler) DeleteBook(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetBooksByAuthorName(c *gin.Context) {

	authorName := c.Query("name")
	if authorName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author name is required"})
		return
	}

	books, err := h.repo.GetBooksByAuthorName(authorName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBooksByCategory(c *gin.Context) {

	categoryName := c.Query("name")
	if categoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category name is required"})
		return
	}

	books, err := h.repo.GetBooksByCategory(categoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}