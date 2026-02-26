package main

import (
	"bookstore/config"
	"bookstore/controller"
	"bookstore/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.InitDB()

	repo := repository.InjectDB(db)
	h := controller.NewHandler(repo)


	r.POST("/books", h.SaveBook)
	r.GET("/books", h.GetBooks)
	r.GET("/books/:id", h.GetBookByID)
	r.PUT("/books/:id", h.UpdateBookByID)
	r.DELETE("/books/:id", h.DeleteBook)

	r.Run(":8080")
}