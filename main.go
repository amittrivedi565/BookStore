package main

import (
	"bookstore/config"
	"bookstore/controller"
	"bookstore/repository"
	"bookstore/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
    defer db.Close()
	repo := repository.InjectDB(db)
	handler := controller.NewHandler(repo)

	r := gin.Default()

	routes.RegisterRoutes(r, handler)

	r.Run(":8080")
}