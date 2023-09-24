package main

import (
	"practice/book_api/cmd/book_api/handlers"
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/logger"
	"practice/book_api/pkg/middleware"
	"practice/book_api/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	bookSvc := service.NewBookService(domain.NewBooksDb())
	bh := handlers.NewBooksHandler(bookSvc)

	router.GET("/books", bh.GetBooks)
	router.GET("/books/:id", bh.GetBookByID)
	router.POST("/books", bh.PostBooks)

	logger.Info("Starting Application...")

	router.Run("localhost:8080")
}
