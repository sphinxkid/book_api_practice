package main

import (
	"practice/book_api/cmd/book_api/handlers"
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	bh := handlers.BookHandler{
		BooksDb: domain.NewBooksDb(),
	}

	router.GET("/books", bh.GetBooks)
	router.GET("/books/:id", bh.GetBookByID)
	router.POST("/books", bh.PostBooks)

	router.Run("localhost:8080")
}
